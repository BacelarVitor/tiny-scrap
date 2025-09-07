// main.go
package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Config struct {
	BaseURL     string
	Start       int
	End         int
	Term        string
	Concurrency int
	Timeout     time.Duration
	ReadCapMB   int
	UserAgent   string
	FullBody    bool // NEW: when true, dump full content in markdown
}

type EntryResult struct {
	Index       int
	URL         string
	Title       string
	Date        string // not extracted in this MVP
	ContentText string
	Matched     bool
	Where       string // "title" or "body"
	Snippet     string
	Err         error
}

func main() {
	cfg := parseFlags()

	// Build job list [start..end]
	jobs := make([]int, 0, cfg.End-cfg.Start+1)
	for i := cfg.Start; i <= cfg.End; i++ {
		jobs = append(jobs, i)
	}

	// Auto concurrency if 0 or negative
	if cfg.Concurrency <= 0 {
		auto := runtime.NumCPU() * 32 // network-bound: many sockets ok
		if auto < 1 {
			auto = 1
		}
		if auto > len(jobs) {
			auto = len(jobs)
		}
		if auto < 1 {
			auto = 1
		}
		cfg.Concurrency = auto
	}

	client := &http.Client{Timeout: cfg.Timeout}

	// Precompile case-insensitive regex for the term
	re, err := regexp.Compile("(?i)" + regexp.QuoteMeta(cfg.Term))
	must(err)

	results := crawlRange(cfg, client, jobs, re)

	// Sort by index and build markdown for matched entries
	sort.Slice(results, func(i, j int) bool { return results[i].Index < results[j].Index })
	md := buildMarkdown(cfg, results)

	outName := fmt.Sprintf("tiny_projects_%s_%d-%d.md", safeFile(cfg.Term), cfg.Start, cfg.End)
	must(os.WriteFile(outName, []byte(md), 0644))
	fmt.Printf("✓ Markdown written: %s\n", outName)

	// Summary
	ok, fail, matches := 0, 0, 0
	firstIdx, firstTitle := -1, ""
	for _, r := range results {
		if r.Err != nil {
			fail++
			continue
		}
		ok++
		if r.Matched {
			matches++
			if firstIdx == -1 || r.Index < firstIdx {
				firstIdx, firstTitle = r.Index, r.Title
			}
		}
	}
	fmt.Printf("Done. OK: %d  FAIL: %d  Matches: %d  (Concurrency=%d)\n", ok, fail, matches, cfg.Concurrency)
	if firstIdx != -1 {
		fmt.Printf("First occurrence at #%d — %s\n", firstIdx, firstTitle)
	} else {
		fmt.Println("No matches found.")
	}
}

// ========================= Worker pool =========================

func crawlRange(cfg Config, client *http.Client, jobs []int, re *regexp.Regexp) []EntryResult {
	jobCh := make(chan int)
	resCh := make(chan EntryResult)

	var wg sync.WaitGroup
	// Workers
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobCh {
				resCh <- processOne(cfg, client, n, re)
			}
		}()
	}

	// Feed jobs
	go func() {
		for _, n := range jobs {
			jobCh <- n
		}
		close(jobCh)
	}()

	// Close resCh when all workers finish
	go func() {
		wg.Wait()
		close(resCh)
	}()

	// Collect results
	results := make([]EntryResult, 0, len(jobs))
	for r := range resCh {
		results = append(results, r)
	}
	return results
}

func processOne(cfg Config, client *http.Client, n int, re *regexp.Regexp) EntryResult {
	u := fmt.Sprintf("%s/%d", strings.TrimRight(cfg.BaseURL, "/"), n)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	body, ct, err := fetch(ctx, client, u, cfg.ReadCapMB, cfg.UserAgent)
	if err != nil {
		return EntryResult{Index: n, URL: u, Err: err}
	}
	if !strings.HasPrefix(strings.ToLower(ct), "text/html") {
		return EntryResult{Index: n, URL: u, Err: fmt.Errorf("unsupported content-type: %s", ct)}
	}

	title, text := parseHTML(body)

	// Match: title → body
	if re.MatchString(title) {
		return EntryResult{
			Index:       n,
			URL:         u,
			Title:       title,
			ContentText: text,
			Matched:     true,
			Where:       "title",
			Snippet:     snippetAround(title, re),
		}
	}
	if re.MatchString(text) {
		return EntryResult{
			Index:       n,
			URL:         u,
			Title:       title,
			ContentText: text,
			Matched:     true,
			Where:       "body",
			Snippet:     snippetAround(text, re),
		}
	}
	return EntryResult{
		Index:       n,
		URL:         u,
		Title:       title,
		ContentText: text,
		Matched:     false,
	}
}

// ========================= HTTP (timeout + read cap) =========================

func fetch(ctx context.Context, client *http.Client, rawURL string, readCapMB int, ua string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, "", err
	}
	if ua != "" {
		req.Header.Set("User-Agent", ua)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp.Header.Get("Content-Type"), fmt.Errorf("status %d", resp.StatusCode)
	}

	// Cap the read size
	capBytes := int64(readCapMB) * 1024 * 1024
	lr := io.LimitReader(resp.Body, capBytes)
	data, err := io.ReadAll(lr)
	if err != nil {
		return nil, resp.Header.Get("Content-Type"), err
	}
	return data, resp.Header.Get("Content-Type"), nil
}

// ========================= HTML parsing =========================

func parseHTML(b []byte) (title string, textMain string) {
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		// Fallback: remove tags crudely
		_, text := stripTagsFallback(b)
		return "", text
	}

	var bt, bx strings.Builder

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			tn := strings.ToLower(n.Data)
			// Skip whole subtrees we don't care about
			if tn == "script" || tn == "style" || tn == "nav" || tn == "footer" {
				return
			}
			if tn == "title" {
				bt.WriteString(textOf(n))
			}
		case html.TextNode:
			t := strings.TrimSpace(n.Data)
			if t != "" {
				bx.WriteString(t)
				bx.WriteByte('\n')
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	title = normalizeSpaces(bt.String())
	textMain = normalizeSpaces(bx.String())
	return
}

func textOf(n *html.Node) string {
	var b strings.Builder
	var rec func(*html.Node)
	rec = func(x *html.Node) {
		if x.Type == html.TextNode {
			b.WriteString(x.Data)
		}
		for c := x.FirstChild; c != nil; c = c.NextSibling {
			rec(c)
		}
	}
	return b.String()
}

func stripTagsFallback(b []byte) (string, string) {
	in := false
	out := make([]rune, 0, len(b))
	for _, r := range string(b) {
		if r == '<' {
			in = true
			continue
		}
		if r == '>' {
			in = false
			out = append(out, ' ')
			continue
		}
		if !in {
			out = append(out, r)
		}
	}
	s := normalizeSpaces(string(out))
	return "", s
}

// ========================= Matching helpers =========================

func snippetAround(s string, re *regexp.Regexp) string {
	loc := re.FindStringIndex(s)
	if loc == nil {
		if len(s) <= 160 {
			return s
		}
		return s[:160]
	}
	start := loc[0]
	const W = 80
	from := max(0, start-W)
	to := min(len(s), start+W)
	return s[from:to]
}

// ========================= Markdown output =========================

func buildMarkdown(cfg Config, results []EntryResult) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Tiny Projects\n\n")
	fmt.Fprintf(&b, "## Entries about: %s\n\n", cfg.Term)

	for _, r := range results {
		if !r.Matched || r.Err != nil {
			continue
		}
		fmt.Fprintf(&b, "### %d — %s\n", r.Index, safeLine(r.Title))
		if r.Date != "" {
			fmt.Fprintf(&b, "*%s*\n", safeLine(r.Date))
		}

		if cfg.FullBody {
			// dump full content
			if strings.TrimSpace(r.ContentText) != "" {
				fmt.Fprintf(&b, "%s\n\n", indent(r.ContentText))
			}
		} else {
			// default: snippet / short preview
			if r.Snippet != "" {
				fmt.Fprintf(&b, "%s\n\n", indent(safeLine(r.Snippet)))
			} else if len(r.ContentText) > 0 {
				preview := r.ContentText
				if len(preview) > 300 {
					preview = preview[:300] + "…"
				}
				fmt.Fprintf(&b, "%s\n\n", indent(preview))
			}
		}

		fmt.Fprintf(&b, "[Link](%s)\n\n---\n\n", r.URL)
	}
	return b.String()
}

func indent(s string) string {
	sc := bufio.NewScanner(strings.NewReader(s))
	var out strings.Builder
	for sc.Scan() {
		out.WriteString(sc.Text())
		out.WriteByte('\n')
	}
	return out.String()
}

// ========================= Utilities =========================

func parseFlags() Config {
	var cfg Config
	flag.StringVar(&cfg.BaseURL, "base", "https://example.com", "Base URL (no trailing slash)")
	flag.IntVar(&cfg.Start, "start", 1, "Start index")
	flag.IntVar(&cfg.End, "end", 201, "End index (inclusive)")
	flag.StringVar(&cfg.Term, "term", "golang", "Search term (case-insensitive)")
	flag.IntVar(&cfg.Concurrency, "concurrency", 0, "Number of parallel workers (0 = auto)")
	flag.DurationVar(&cfg.Timeout, "timeout", 8*time.Second, "Per-request timeout")
	flag.IntVar(&cfg.ReadCapMB, "readcap", 4, "Read cap per response (MB)")
	flag.StringVar(&cfg.UserAgent, "ua", "TinyScrap/0.1 (+https://github.com/you)", "User-Agent")
	flag.BoolVar(&cfg.FullBody, "fullbody", false, "If true, dump full page content in Markdown")
	flag.Parse()

	cfg.BaseURL = strings.TrimSpace(cfg.BaseURL)
	if _, err := url.ParseRequestURI(cfg.BaseURL); err != nil {
		fmt.Fprintf(os.Stderr, "invalid base url: %v\n", err)
		os.Exit(2)
	}
	if cfg.Start > cfg.End {
		cfg.Start, cfg.End = cfg.End, cfg.Start
	}
	if cfg.ReadCapMB < 1 {
		cfg.ReadCapMB = 1
	}
	return cfg
}

func safeFile(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r
		case r >= '0' && r <= '9':
			return r
		case r == '-' || r == '_':
			return r
		default:
			return '-'
		}
	}, s)
	return filepath.Clean(s)
}

func safeLine(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\n", " "))
}

func normalizeSpaces(s string) string {
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal:", err)
		os.Exit(1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
