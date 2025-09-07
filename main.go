
/* Fuck
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	var lastEntry int
	baseUrl := "https://daily.tinyprojects.dev"
	term := "prompt"

	html := getHtml(baseUrl)
	re := regexp.MustCompile(`<a\s+href="(/\d+)"`)

	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		length := len(match[1])
		value := match[1]
		lastEntry, err = strconv.Atoi(value[1:length])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < lastEntry; i++ {
			searchTerm(i)
		}
	} else {
		fmt.Println("No match found")
	}
}

func searchTerm(i int) {
	html := getHtml()
	re := regexp.MustCompile(`<a\s+href="(/\d+)"`)

	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		length := len(match[1])
		value := match[1]
		lastEntry, err = strconv.Atoi(value[1:length])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < lastEntry; i++ {
			searchTerm(i)
		}
	} else {
		fmt.Println("No match found")
	}
}

func getHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
*/

// main.go
