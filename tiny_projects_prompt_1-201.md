# Tiny Projects

## Entries about: prompt

### 3 —

3 - First Users | Tiny Projects 3. First Users Wednesday, September 15th 2021 Yesterday was a great success! My aim was to get 1 user for Paper Website. But as of this morning I have 9 (all free for now). I did 2 things. Firstly, I whipped up a short TikTok video about Paper Website and posted it online around 5 pm. It's currently sitting on 238 views and 4 likes. Not a viral banger but okay. Secondly, just as I was about to switch off for the evening, I noticed there was a trending post on Hacker News about tips for journaling. I jumped on it early and posted a link to Paper Website . Basically, it worked very well - it's the top comment, and is still bringing in 60 - 80 website visits an hour. All the comments are really nice too! (This was not the case with Mailoji..). Although I have 9 users, none of them have created a blog post yet. I think this is because there's so many hurdles like creating your blog post, downloading the camera app, editing etc. etc. Maybe these people signed up and were planning to use it later, it does make sense. I think I'm going to create some sort of automated email system that prompts people to write each morning. If people regularly start using Paper website then that's where things will get interesting. My goal for the end of the day is to add these email notifications and get at least one user to write a blog post. Here's my view by the way! << Next Prev >>

[Link](https://daily.tinyprojects.dev/3)

---

### 4 —

4 - Three X'd | Tiny Projects 4. Three X'd Thursday September 16th 2021 Dropping comments on Hacker News comments is a great way to get users! I now have 28, so 3Xed over night! I've also had some users try handwriting to website conversion. I noticed that the results were mixed, so I spent some time yesterday improving the main algorithm, although I think that made it worse... something I will finalize today. So far 1000 people have visited the homepage, with 28 users, that's a conversion rate of 2.8% - I thought this would be higher given it's currently a free product, and feedback is good, so I want to figure out what's up. I also need to create some email prompts to get people writing, as the users trying it out are submitting test examples at the moment. Lots to do! << Next Prev >>

[Link](https://daily.tinyprojects.dev/4)

---

### 26 —

26 - Being productive whilst waiting to launch | Tiny Projects 26. Being productive whilst waiting to launch Monday October 18th 2021 Good Morning! I had a pretty chill weekend catching up on some work and playing a lot of New World - a game I'm still really enjoying. I feel like I'm in a bit of a waiting period until my custom notebooks arrive for Paper Website, the new tiny project that I'm building. My plan is to get notebooks then launch as soon as possible, as I've been waiting too long now and I just want to get it released. Everything now feels good on the Paper to Website conversion. I deep dived into the algorithm and altered the prompt I'm using for GPT-3, which is the bit of text that contains instructions and examples to tell GPT-3 what to do. I basically added more examples and clearer instructions. To test turning handwriting into a website, there is conveniently a subreddit called /r/handwriting. There's so many different examples and styles to choose from, but Paper Website can handle most of them. Even the bad handwriting of a 3 year old can be turned into a website - how cool is that?! Something else I added recently was a crop feature, which has improved using Paper Website 100%. Basically, when you take a photo of your page, you need to make sure no other background text is visible, otherwise it gets picked up and confuses the system. This is usually the edge of another page, or a screen or keyboard in the background of the photo. Previously this was quite annoying, as you needed very steady hands to take a perfect picture that didn't pick up any background text. Now through there is a cropping feature, so you can easily submit only the page you want to convert. With little engineering to do on my tiny projects this week I think I'll just see what happens when I fully set my mind on marketing. feel usually when I market, I only speed around 30 minutes focused doing it. What happens if I do it for 4 hours? What about a whole day? I'm interested to find out. My hand hurts, this was a longer post than I thought! << Next Prev >>

[Link](https://daily.tinyprojects.dev/26)

---

### 136 —

136 - Introducing PromptBase, my next tiny project! | Tiny Projects 136. Introducing PromptBase, my next tiny project! Thursday, June 2nd, 2022 Here it is: <https://promptbase.com> PromptBase is a website for buying and selling prompts for AI models like GPT-3. Prompts are basically the instructions you use to tell an AI like GPT-3 what to do, it's like code but it can do really powerful things. This idea was born out of Paper Website, where I needed a prompt to fix spelling mistakes in OCR text. I couldn't find one online, but I would have paid good money for one. Instead I had to spend hours making my own and learning "Prompt Engineering". (this is not the OCR prompt I'm using, just an example of a 1st version I made) At the same time, a few people have asked about the prompt I made, and I thought it would be cool to make some passive income selling these magic text files. Example of a Prompt for sale on PromptBase Good prompts produce better results, and are cheap to run in terms of API costs, so if you have a good prompt it can save someone a lot of money. On top of this, my bet is that programming is going to change a lot in the next decade. I think I'll spend less time writing Javascript, and more time writing prompts. There's been occasions at my regular job where I've used prompts instead of code and saved weeks of time. Although this only happens maybe 1% of the time currently, it's increasing. Anyone can sell a prompt on PromptBase, and on each sale I take a fee, a bit like an app store. Currently the split is 80% to the prompt creator (all handled via Stripe). I'm also experimenting with a custom prompts page, where someone can request a bespoke prompt, so there's a few ways this idea can be monetized. My next steps are to get prompt sellers onboarded and create a well stocked marketplace (currently there's just 2 prompts for sale!). I also think this website is going to live and die on SEO, so I'm going to start optimising that, and perhaps experimenting with some Google Ads. I'll be trying to market and improve the site this month, then will launch on ProductHunt/Twitter and write a full blog post about it afterwards. I'll share updates as I go on this daily blog in the meantime. If you have any feedback on the idea or the website, I'd love to hear it! << Next Prev >>

[Link](https://daily.tinyprojects.dev/136)

---

### 137 —

137 - First Prompt Sale | Tiny Projects 137. First Prompt Sale Monday, June 6th, 2022 One of my goals since moving to Australia was to learn how to surf. This weekend, I went to a surf lesson and managed to catch some waves! I suck, but it was so fun. I've now bought a board to try and improve. In exciting news, PromptBase got its first sale! One OCR corrector prompt for $9.99 - I think this might be the first documented prompt sale ever. So cool. I have three goals for PromptBase in June: Add more prompts Improve SEO Improve the look of the website Off to a great start. << Next Prev >>

[Link](https://daily.tinyprojects.dev/137)

---

### 138 —

138 - Improving SEO | Tiny Projects 138. Improving SEO Tuesday, June 7th, 2022 I want to improve my SEO on PromptBase. The goal is, if somebody types "GPT-3 prompt for x" into google they should see a PromptBase link. The website is built using Angular, with client-side rendering; probably the worst way to build an SEO-optimised site. To make it so web crawlers can see my meta tags and page content, I need to make the website instead be server-side rendered. Fortunately, there's something called "Angular Universal" that can do this, which doesn't look too hard to implement. I'll report back on how it goes. << Next Prev >>

[Link](https://daily.tinyprojects.dev/138)

---

### 140 —

140 - I bought 275 GPT-3 prompts | Tiny Projects 140. I bought 275 GPT-3 Prompts Tuesday, June 14th, 2022 I feel with every project I build, I splurge money on something - with PromptBase it's no different. One of my goals in June was to add more prompts for sale on the website, to increase my current library size of a measly 2 prompts. My goodness, has that number now increased. Yesterday I purchased 275 prompts off a prompt engineer for $1000. That's a 138 X increase! Mission complete, I would say. Now I need to go through each one, process them, and list them for sale on PromptBase. On top of this I still need to sort SEO and finish writing my article (thanks to everyone who replied by the way!) Busy busy. << Next Prev >>

[Link](https://daily.tinyprojects.dev/140)

---

### 141 —

141 - 4 days of code hell | Tiny Projects 141. 4 days of code hell Friday, June 17th, 2022 4 days. (Well, mornings) That's how long it took me to get Angular Universal Server-Side Rendering working (it's as fun as it sounds). Holy Christ. I think it has to be one of the most fiddly, awkward bits of programming I've ever done. It works at least, now my links look pretty and my SEO should be way better. Here's a before and after example linking to the exact same OCR text corrector prompt webpage: Although this was a nightmare to program, I never have to do this again. It's another tool in the tiny projects toolbelt that I can just keep re-using in every project. Just like authentication, payments, emails, and all the other boring but necessary things that have come before it. Now, I have 275 prompts to upload... << Next Prev >>

[Link](https://daily.tinyprojects.dev/141)

---

### 143 —

143 - Very Busy | Tiny Projects 143. Very Busy Tuesday, June 28th, 2022 Sorry for the lack of updates! I have been extremely busy writing, so there's not been a whole lot going on. This article is almost done, it's shaping up really nicely now. I've spoken with and interviewed some really awesome people: Ben Issen, Daniel Vasallo, Tony Dinh, Ben Awad - my next goal is to speak with the guy who made Wordle. I should be finished in a couple of days, then I'll go back to my neglected tiny projects. Poor PromptBase. << Next Prev >>

[Link](https://daily.tinyprojects.dev/143)

---

### 144 —

144 - How much $ I made in June | Tiny Projects 144. How much $ I made in June Friday, July 1st, 2022 Halfway through the year already, crikey. Here's how much my tiny projects made in June. Mailoji did £150.81, way down from last month's spike, but expected. Paper Website was pleasantly up, at £282.78. I've noticed a bit of growth here since launching the new app. One Item Store is still at £0. Yeah.. I need to give it some more love. PromptBase comes in at £19.56! Although I'm pretty sure all these sales were because of this blog - this is still really cool. In total that comes to £453.15 ($550.55) I am so looking forward to going back to working on these projects in July. I've spent a lot of time writing the past couple of months, not enough time building. << Next Prev >>

[Link](https://daily.tinyprojects.dev/144)

---

### 145 —

145 - Paper Breakthrough | Tiny Projects Paper Breakthrough Monday, July 4th, 2022 The handwriting recognition on Paper Website is good, but I want it to be excellent. I've used the same prompt system since I launched it, however, this weekend when I was working on PromptBase, I discovered something called GPT-3 Fine Tuned Models. Basically, it lets you train the AI on lots of examples, instead of squeezing just a few into a tiny prompt. I have nearly 150 blog posts handwritten with text corrections now, so I thought - why not feed them into GPT-3 and see what happens? $10 and 20 minutes of training later, my fine tuned model was created - and it was working really well. Then, I had another brainwave - not only do I have 150 examples of OCR handwriting corrections, but I've also got 150 examples of turning that corrected text into HTML. E.g. making titles be h1, making the post date italic, making lists have bullet points, adding links to things. Previously, everything just got wrapped in a p tag and you'd style the page yourself. Can you do this automatically? $15 later (HTML is larger to train on), I had my new HTML Fine Tuned Model. I fed in some paper blog posts and they were being automatically styled - it was magic. I'm actually really excited by this, it makes me want to push with Paper Website for a bit. I've made this new engine live for all Paper Website users, this is my first post using it. If you're a user, I'd love your feedback. << Next Prev >>

[Link](https://daily.tinyprojects.dev/145)

---

### 147 —

147 - New Plan | Tiny Projects New Plan Thursday, July 14th, 2022 I've decided to focus my tiny projects efforts on Paper Website for a bit, at least for the next couple of months - and I'm excited by this! For now, this means I'm going to put PromptBase on the backburner. I want to try and grow Paper Website, which, if you scroll all the way back to the start of this blog, is what I was trying to do all along. There's two problems right now which I need to figure out. 1. People don't know what to write about, but - this is because my audience is so broad. I need to find an initial niche - e.g. travel journalers, because they know what to write about: their trip. 2. Getting people to write their first pages and set their website up, and making this as easy as possible. Most people who churn never create their first page. If I can get the user to create a couple of pages and share it with someone, I'm confident they'll stick around. I tweeted in January Paper Website was at $10K ARR, and although new customers have come through the door - it's still at $10K ARR. Hopefully this is the point where that starts to change. << Next Prev >>

[Link](https://daily.tinyprojects.dev/147)

---

### 153 —

153 - A Big Pivot | Tiny Projects A Big Pivot Wednesday, July 27th, 2022 A lot has happened since my last post - long story short, I think I've managed to get PromptBase working after almost binning the idea. I'll rewind. On Saturday morning, I got an email saying I had finally got access to DALL-E , OpenAI's new model for generating images from text. DALL-E uses prompts like GPT-3 to get it to generate images, e.g. generate an image of "a lego version of Homer Simpson in 8K UHD on a white background". PromptBase was already set up and ready to sell GPT-3 prompts, so I wondered: could I change PromptBase to also sell quality DALL-E image prompts? It took about 5 hours of work, but I got PromptBase into a good state where someone could list a DALL-E prompt for sale. Next step was to create some killer prompts to sell. Here are the images generated by some of my best ones: Clay Emojis and Polygon Animal prompt output images I started posting these outputs on Reddit, Twitter, Instagram and Discord, then dropping links to the prompts on PromptBase. Then, side story, the past couple of days I've been in Brisbane - and whilst I was walking about looking at Koalas, I kept getting notifications about PromptBase sales - it was so cool! Damn, koalas are cute I've also got a queue of people lining up to start selling their own prompts, which I need to review. Back in the saddle now and excited to see where this goes. << Next Prev >>

[Link](https://daily.tinyprojects.dev/153)

---

### 154 —

154 - Reviewing Prompts | Tiny Projects Reviewing Prompts Friday, July 29th, 2022 Having a lot of fun working on PromptBase - it's very exciting. My little marketplace has grown to 18 prompts, and the best thing is there's now people submitting their own prompts as well as me creating them. I review each prompt submitted, but it was getting so slow because I was having to do it all through editing fields in a database - so I built a quick and dirty tool to make reviewing prompts easier. Here's a couple of the latest prompts I've created: shoe product shots and tiny planets . << Next Prev >>

[Link](https://daily.tinyprojects.dev/154)

---

### 155 —

155 - PromptBase in TechCrunch! | Tiny Projects PromptBase in TechCrunch! Monday, August 1st, 2022 Here is the article. My goodness - this is getting slightly out of control. I got an email from a reporter asking some questions, but I never expected a full article on PromptBase (also, this is my second TC feature in two weeks..!). Without a doubt, this is going to be the craziest Tiny Project story yet. << Next Prev >>

[Link](https://daily.tinyprojects.dev/155)

---

### 156 —

156 - Triple Threat | Tiny Projects Triple Threat Wednesday, August 3rd, 2022 PromptBase is still growing well - I don't have time to code any features, I just want to get people onboarded and selling - now up to 57 prompts. What's cool is people are now sharing that they're making money from their prompts, which is bringing more people in the door. I remember seeing this tweet from Arvid Kahl a while back: PromptBase ticks all these boxes - it's a triple threat. << Next Prev >>

[Link](https://daily.tinyprojects.dev/156)

---

### 157 —

157 - Pumping Out Thumbnails | Tiny Projects Pumping Out Thumbnails Friday August 5th, 2022 For the past 5 days I've woken up and spent around 3 hours each morning in Photoshop creating thumbnails like these for all the prompts being submitted. I kinda enjoy it, I just listen to podcasts and drink coffee at the same time - it's very chill. I'm doing this for two reasons: 1. Makes the store look nicer - I can't trust users to make these (yet). 2. Hopefully makes the prompts more appealing to buy, meaning prompt sellers generate more income and write more prompts. Fun fact: Back in the golden era of 2010, I used to play a game called RuneScape a lot - so much so, that me and a couple of friends started making videos about it and posting them to YouTube (we hit about 6k subscribers!). I probably made around 100 video thumbnails back then, it's very funny that this random skill has come into good use again. Now at 75 prompts listed! << Next Prev >>

[Link](https://daily.tinyprojects.dev/157)

---

### 158 —

158 - A Few New Features | Tiny Projects A Few New Features Monday, August 8th, 2022 Added some cool features to PromptBase over the weekend - mainly for sellers. First, there's now more stats on your prompts, including views and sales. Second, there's now a "More from this prompt engineer" section under each prompt page. Third is sections for new prompts , and a top chart of the most popular prompts. Eventually I'll have different categories like "Logos", "Icons" and "3D". It's coming along nicely - but that's enough coding for now, back to uploading prompts and marketing. By the way - that article I was writing about building tiny projects is out tomorrow, it'll be on a16z's Future website! I'll drop the link here too though. << Next Prev >>

[Link](https://daily.tinyprojects.dev/158)

---

### 159 —

159 - Rest, Rickshaws and a16z | Tiny Projects Rest, Rickshaws & a16z Friday, August 12th, 2022 Took a few days off to go lie on a beach and attempt to surf - it was much needed. I've been pulling some very long days the past few weeks. PromptBase kept going fine - I scheduled a load of prompts to be uploaded each day along with new social media posts. Tuesday was the biggest day of sales yet. In a few weeks time I'm doing something called "The Rickshaw Run", where me and a couple of friends are driving a rickshaw from the south of India to the north over 2 weeks. I need to figure out how to keep the website going over that period, as I doubt I'll have much time. I might even hire someone. The article I wrote for a16z Future also came out on Tuesday, it's called "Why developers are building so many side projects" , give it a read if you fancy! << Next Prev >>

[Link](https://daily.tinyprojects.dev/159)

---

### 160 —

160 - Ripping Off Netflix | Tiny Projects Ripping Off Netflix Monday, August 15th, 2022 There's now about 150 prompts on PromptBase , and it had got to the point where it was becoming actually hard to find a prompt. Basically, every prompt just appeared in one big list, which was quite overwhelming and bad user experience. So, this weekend I went through every prompt and added tags to them, e.g. #3d #logos #icons . Then, I built these horizontally scrolling prompt collections, and used them to house categories. I took a lot of inspiration from the Netflix homepage, but I'm really happy with how it turned out. Tagging also lets me do cool things like, if a user is viewing a prompt relating to #cute #2d #robots , suggesting other similar prompts below it. I think these categories also makes it a bit clearer with what I'm going for with PromptBase. I think the most exciting thing about prompts is they have utility (e.g. generating quality website icons), as apposed to just being about generating art like "Homer Simpson on a beach in watercolour". Anyway, back to photoshop. << Next Prev >>

[Link](https://daily.tinyprojects.dev/160)

---

### 161 —

161 - Battle of the Bandwidths | Tiny Projects Battle of the Bandwidths Wednesday, August 17th, 2020 I decided to peek at my Firebase project cost for PromptBase the other day and noticed 2 things: 1. It was way higher than I thought it would be. 2. It was shooting up fast. For context, it costs me about £5/month to run Mailoji. The main cost was from "cloud storage bandwidth". When a user uploads a prompt, they upload some example images with it to Firebase storage. Then, when a user is viewing the store page for that prompt, the images served on that page are also from cloud storage. The problem was, I was serving the raw images uploaded by the user from Dall-E, which are 1024 × 1024 and about 1.5 MB each. On a store page with 9 example images, that's over 10 MB of bandwidth per page view! Not good. Fortunately, Firebase have this great "Image resizing" extension that I could just turn on, and now my images are resized and served at a much smaller size of 600 x 600, with a lower bandwidth. I know I shouldn't be optimising for costs at this stage, but it felt like an important one over the long haul for image-heavy PromptBase. << Next Prev >>

[Link](https://daily.tinyprojects.dev/161)

---

### 162 —

162 - Dollar Mail | Tiny Projects Dollar Mail Friday, August 19th, 2022 My biggest goal with PromptBase right now is growth. However, it's a marketplace, which is notoriously hard - I need new customers, and I also need new prompts coming through the door. Today I added sales notification emails for whenever you make a sale. Users could previously passively check sales in their dashboard - but they were never actively alerted when they made a sale. Yes, these emails are useful - but my sneaky ulterior motive is really to constantly remind users they're making money, and therefore encourage them to make more prompts. Now at 207 prompts uploaded. << Next Prev >>

[Link](https://daily.tinyprojects.dev/162)

---

### 163 —

163 - Whale Alert | Tiny Projects Whale Alert Monday, August 22nd, 2022 This weekend I went and watched some Whales - they were awesome! I've pretty much ticked off every major Australian animal at this point, except maybe a snake (but all good on that one). I also spent some time trying to automate away PromptBase tasks that were taking too long. A basic example is creating the URL slugs for each prompt page, e.g. e.g. <https://promptbase.com/prompt/> rpg-item-icons Previously I would manually type these slugs, e.g. if a prompt was called "Retro Synthwave Animals" I'd hand type out an appropriate slug, like "retro-synthwave-animals". Sure, this doesn't take too long - but it was adding up over time. Now I just click a button (same with tags). It's so much faster! << Next Prev >>

[Link](https://daily.tinyprojects.dev/163)

---

### 164 —

164 - Birth of a Job | Tiny Projects Birth of a Job Wednesday, August 24th, 2022 I've hired someone to review prompts for PromptBase whilst I'm driving across India in a rickshaw. They are... my brother! We hopped on a call a couple of days ago and I gave him a crash course in prompt reviewing. He starts next week. So, what does a prompt reviewer do? Mainly, they test Dall-E, Midjourney and GPT-3 prompts to make sure they work well (and are not a scam, oh boy - that's a whole other story). When a good prompt is submitted, the prompt reviewer makes it look good for the website, e.g. making a thumbnail, setting the price, writing a good description. The concept of a "prompt engineer" has been around for a while, but I'd hedge my bets and say this is the first documented "prompt reviewer" to exist. What do you want to be when you grow up, son? A prompt reviewer. << Next Prev >>

[Link](https://daily.tinyprojects.dev/164)

---

### 165 —

165 - 1,000 Users | Tiny Projects 1,000 Users Friday, August 26th, 2022 1,259 actually! It's been exactly 1 month since I pivoted PromptBase , it was pretty much dead at that point and had maybe 10 users. This month has been something else. I've not really been sharing figures, but this is a cool milestone to celebrate. Let's hope it keeps up. << Next Prev >>

[Link](https://daily.tinyprojects.dev/165)

---

### 166 —

166 - Going Vertical | Tiny Projects Going Vertical Monday, August 29th, 2022 Something very cool happened last week - a new AI art generator was released called "Stable Diffusion". It's just like DALL-E, except it's open source, and you can run it yourself. So, I had an idea - could I somehow get Stable Diffusion into PromptBase? After a weekend of tinkering, I did it! You can now generate images and craft prompts directly in PromptBase, then list those prompts on the marketplace with a couple of clicks. This opens up Prompt Engineering to anyone - if you don't have access to DALL-E yet, you can literally generate images now on PromptBase for free (well, you get 10 free generations/day). Owning the whole stack has opened so many possibilities. << Next Prev >>

[Link](https://daily.tinyprojects.dev/166)

---

### 167 —

167 - Eyeballs | Tiny Projects Eyeballs Wednesday, August 31st, 2022 More PromptBase features - one I'd been lacking for ages is search filters for categories and AI Models. This should hopefully make things easier for a buyer, and surface prompts that otherwise wouldn't get enough views. Speaking of which - I've been collecting "views" data for each prompt behind the scenes for a few weeks now. Previously, the only people who could see the view numbers were me, to analyse which prompts are popular, and sellers to track how their prompts are doing. This morning I decided to make this view number public on each prompt page to anyone. I deliberated about this. Sharing view numbers felt odd, like perhaps it would expose PromptBase in some way. But, the more I thought about it, the pros seemed to outweigh the cons. Personally, I think if I saw that a lot of people had viewed a page, it would give me more trust, and give me the sense that there's activity on the website. The same way you might be more excited/trust a YouTube video more if it had 10M views instead of not knowing the view count. I'm going to monitor this - I can always change it back. But, interested to hear your thoughts - show views or hide views? << Next Prev >>

[Link](https://daily.tinyprojects.dev/167)

---

### 168 —

168 - PromptBase in The Verge! | Tiny Projects PromptBase in The Verge! Monday, September 5th, 2022 Big article from The Verge about PromptBase came out over the weekend - it was an interview with a user selling their prompts on the marketplace - really good! Numbers have all skyrocketed: sales, prompts uploaded, page visits. Last week I reported having 1,000 registered users, now that number is about 2,500. Every project I've launched previously kinda fizzles out after a few months, e.g. they get a huge spike after a Hacker News/Product Hunt launch, then it quietens down. Each day I expect this to happen with PromptBase - in fact, I still half think it probably will when the current text-to-image hype dies down a bit. Right now though, it's not showing any signs of slowing down. << Next Prev >>

[Link](https://daily.tinyprojects.dev/168)

---

### 169 —

169 - Rickshaw Run | Tiny Projects Rickshaw Run Wednesday, September 7th, 2022 Currently writing this in Kuala Lumpur airport! I'm heading to India so I can drive 3,000km across it in a rickshaw for "The Rickshaw Run". Not 100% sure how often I'll be able to update this blog over the next few weeks. I have no idea where I'm going or if there'll even be wifi. Hopefully PromptBase will survive. Things are still very busy, but fortunately my brother is now pretty pro at reviewing and uploading prompts to the site, so it should be fine. See you on the other side! << Next Prev >>

[Link](https://daily.tinyprojects.dev/169)

---

### 170 —

170 - Back From India | Tiny Projects Back From India Monday, October 3rd, 2022 Back from India! Had a great time - way too much happened to fit into a daily blog. In total we drove 3,700km in our rickshaw from the south (Kochi) to the north (Jaisalmer), swinging by the Taj Mahal along the way. With a top speed of 55 kmph, it took about 2 weeks of solid driving. Crazy roads: cows, dogs, potholes, camels and even elephants everywhere that you need to dodge. Only broke down once - Indian people are the most friendly people I've ever met so had lots of help to fix. Our rickshaw was designed by a Dall-E prompt! which you can of course buy on PromptBase. Leading onto PromptBase - the entire time I was gone it grew like crazy - it was actually pretty stressful not being at a keyboard whilst it was happening. I'll report back in the next post with some stats. << Next Prev >>

[Link](https://daily.tinyprojects.dev/170)

---

### 171 —

171 - 7,000 Users | Tiny Projects 7,000 Users Wednesday, October 5th, 2022 PromptBase had some fantastic growth over the past 3 weeks - almost 5,000 new registered accounts were added! The marketplace as a whole is feeling healthier, with a good number of buyers and sellers joining each day. I have a few goals this month: 1. Review times for prompts are too long due to volume - it's gone from 1 day to 10 days - I'm exploring ideas to speed this up, no one wants to wait that long. 2. More visibility for Prompt Engineers - currently sellers are pretty anonymous on PromptBase, I'd like to add profiles and usernames to change this. 3. Improve a bunch of other site problems, like search to make it easier to find prompts. Excited to get to work! << Next Prev >>

[Link](https://daily.tinyprojects.dev/171)

---

### 172 —

172 - Prompt Profiles | Tiny Projects Prompt Profiles Monday, October 10th, 2022 Felt very rusty with coding and project building last week after so much time off - but slowly getting back into it. Yesterday I added a user profiles to PromptBase . I'm really happy with how they turned out! As a seller, you now have a space for your prompts, and a place to share links to your socials and website. Here's my page: <https://promptbase.com/profile/ben> I also decided to show total sales, views, and a "PromptBase Rank" for the top 100 creators. My main goal with this feature is to make it clearer that there's real people making these prompts, instead of just a faceless website. It's also step #1 for some other future exciting features - more to come! << Next Prev >>

[Link](https://daily.tinyprojects.dev/172)

---

### 173 —

173 - One HTML Tag Saved Me £300 | Tiny Projects One HTML Tag saved me £300 Wednesday, October 12th, 2022 Isn't that a clickbaity title! No joke though. One of my biggest expenses with PromptBase is image bandwidth - there's images everywhere on the website. Every time someone views a PromptBase image in their browser, I'm paying some small amount in data bandwidth to send the bits and bytes to them. One image obviously doesn't cost too much, but there's been days on PromptBase where over 100 GB of image data is being streamed to all the website visitors. Previously, I thought I'd optimised this by serving smaller images, but my costs were still really high, last month it was over £400 (my other projects are around £10/mo). I looked into it again, this time paying closer attention to exactly how much image data was being loaded in on the marketplace page. Over 80 MB. Christ! The problem: I was loading in hundreds of images off screen (i.e. down the page when you scroll). I needed to lazy load these images, so they only loaded in when they were near or on screen. I've done lazy loading before using various libraries. But it's a bit of a pain, so I wasn't thrilled. You could say I was feeling pretty lazy about it. But, as I started to delve into it, I came across a newish attribute for the <img> HTML element that apparently released a few years back: <img src="" loading="lazy" /> I thought it was too good to be true. But, I added the lazy loading tag and retried my page load experiment. The image data loading in dropped from 80 MB to 6 MB - over a 10X reduction. I've had this in for a few days and am already seeing results - if I'd had it in last month it probably would have saved me £300 (about $300 at today's exchange rate..) Native lazy loading - pretty cool. << Next Prev >>

[Link](https://daily.tinyprojects.dev/173)

---

### 174 —

174 - High Speed Prompt Reviews | Tiny Projects High Prompt Speed Reviews Friday, October 14th, 2022 The past few mornings I've been working on an exciting new feature for PromptBase - I should hopefully be able to launch it next week! The only reason I've been able to work on new features is because I've massively sped up the "Prompt reviewing" process that was taking up so much time. Even though I had my brother helping me with reviewing when I was away in India, I still returned to over 1,000 pending prompts; some of which had been waiting 2 weeks for reviews. Now, I've reduced prompt reviewing to testing out the prompt, and basically clicking a "no" button if it doesn't work, or a "yes" button if it's good. The latter generates tags, formats the title, schedules it, and even automatically generates a thumbnail - no more photoshop required! Auto-thumbnail on an e-sports logo prompt 2 months ago, reviewing a prompt took 15 minutes. Now, it takes 30 seconds! << Next Prev >>

[Link](https://daily.tinyprojects.dev/174)

---

### 175 —

175 - Tiny Social Growth | Tiny Projects Tiny Social Growth Monday, October 17th, 2022 I've been pretty quiet on my own Twitter account the past few months whilst I've been working away on PromptBase (I haven't even "twitter-launched" PromptBase yet..!). But, I have been incredibly active on the PromptBase Twitter and Instagram account. Each day I post a 3x3 collage of images generated by some of the best prompts sold by users. The insta feed. It's a real highlight of my day doing this. It feels a lot more organic compared to my previous tiny project social media attempts , and I'm actually seeing some small results. At the time of writing this there's 442 followers on Twitter and 260 on Instagram. This might finally be a marketing channel I actually enjoy! << Next Prev >>

[Link](https://daily.tinyprojects.dev/175)

---

### 176 —

176 - Still Building | Tiny Projects Still Building Wednesday, October 19th, 2022 Still building my big new PromptBase feature. It's really exciting, and could potentially be a game-changer in terms of revenue and the direction of the site. Few more days to go - sorry for the boring update! << Next Prev >>

[Link](https://daily.tinyprojects.dev/176)

---

### 177 —

177 - Soft Launching | Tiny Projects Soft Launching Friday, October 28th, 2022 Okay, this new feature took longer than a few more days - sorry for the lack of updates. I've just spent the past week coding. But - it's done! I'm going to soft-launch the feature this weekend then announce it / write about it next week. It's basically a whole new project within a project. The Financial Times did a great article yesterday about AI art, including a PromptBase shoutout. I couldn't agree more with Miller. << Next Prev >>

[Link](https://daily.tinyprojects.dev/177)

---

### 178 —

178 - Hire a Prompt Engineer | Tiny Projects Hire a Prompt Engineer Wednesday, November 2nd, 2022 Just launched! A brand new section in PromptBase: Hire a Prompt Engineer: <https://promptbase.com/hire> Instead of buying a standalone prompt file, you can commission a talented prompt engineer on PromptBase to create a prompt for you. You can find a Prompt Engineer you like via their profile. Then start a chat with them and pre-pay for your prompt, then the prompt engineer creates your prompt for you. It's a brand new income stream for a prompt engineer - and potentially the first place on the internet to hire one. Very excited to see how it goes. I've had quite a few emails from both buyers and sellers requesting this feature - so I'm confident the demand is there. Building this was challenging but pretty fun - I basically created my own WhatsApp clone with integrated payments for the chat window. The best thing was - I launched this hire page by messaging the 100 top sellers through my own chat feature. It was so awesome to speak with some of these insanely talented people, and a great way to get feedback. I'll report back on how it goes, and hopefully return to more regular updates on this blog! << Next Prev >>

[Link](https://daily.tinyprojects.dev/178)

---

### 179 —

179 - First Jobs | Tiny Projects First Jobs Monday, November 7th, 2022 My internal micro job-tracker dashboard First few jobs have come through and been paid for on PromptBase - very exciting! I've purposely left the whole system quite open to start like a big experiment: just a chat window and payments. I want to see what naturally happens. Already I'm seeing areas where much more clarification is needed, like turnaround time, number of revisions, and setting expectations for final prompts and produced images. Going okay so far though! << Next Prev >>

[Link](https://daily.tinyprojects.dev/179)

---

### 180 —

180 - New Prompt Problems | Tiny Projects New Prompt Problems Tuesday, November 8th, 2022 I have a really interesting problem on PromptBase. When a new prompt goes live, it appears in the "newest" section. Currently it gets 3-4 hours of exposure there before it goes horizontally off-screen because even newer prompts have taken its place. How new prompts flow off the page on PromptBase Every category across PromptBase is ordered by most popular of all time, to showcase the best prompts to a potential buyer. All PromptBase categories on the marketplace are ordered by most popular. However, this means after a prompt has left the Newest Prompts section, it's almost impossible for that prompt to get exposure, even when it might be really good! It's lost to the oblivion unless you search for it. This is really bad, because if a user creating good prompts isn't getting sales, they're not going to keep submitting prompts, and the site will die. My current solution is this new "trending" section, which is all across the site now. New trending prompts section It uses a modified version of the Hacker News algorithm for surfacing new prompts that receive lots of views and sales. The Hacker New algorithm. After that, I've built a sort-of funnel. There's a section for top prompts of the week, then top prompts of the month that good new prompts should hopefully flow through to increase their exposure. I think it's sort of working. A minor benefit also is the homepage feels less stale - prompts are constantly changing so a user might check back more often (like checking the front page of Reddit). I'm tempted to change the whole marketplace page to use the trending algorithm, but there's a risk it will show worse prompts to a potential buyer which will reduce demand, and then I'm back in the same position of prompts not getting sales. But, at the same time, are the most popular prompts only popular because there's not a trending algorithm everywhere? It's a real chicken-and-egg problem! << Next Prev >>

[Link](https://daily.tinyprojects.dev/180)

---

### 181 —

181 - Double Speed | Tiny Projects Double Speed Wednesday, November 9th, 2022 Straight after I wrote the daily blog post yesterday , I decided to switch the entire front page of the marketplace from being ranked by "top of all time" to using the new trending algorithm. This is a big change, but I think it's more forward-thinking to keep the supply side of the marketplace healthy. Supply, in theory, is also a good lever for increasing demand too (people tell their friends they sold a prompt, etc.). I'm now experimenting with doubling the speed at which new prompts are published to the site from 24 to 48/day. Also, I'm doubling the frequency of social posts from one to two a day. There's now 2192 live prompts uploaded! << Next Prev >>

[Link](https://daily.tinyprojects.dev/181)

---

### 182 —

182 - Spam to Steak | Tiny Projects Spam to Steak Thursday, November 10th, 2022 The new in-website messaging feature for PromptBase is working surprisingly well! Each day I usually have one or two messages from sellers through it, which is super useful for feedback - and I think is sort of helping to build a tiny community around PromptBase. I'm genuinely considering throwing every seller into a single chatroom on there and making my own in-website Discord! It's not all been plain sailing though - the first day of it launching someone mass-spammed 50 users with a strange job request, trying to pull them outside of the website for payments etc. The whole experience was pretty annoying, but it did give me the idea that I could probably charge companies to post AI-related job ads on PromptBase and have prompt engineers come to them instead. Maybe a valuable insight from a bad situation! << Next Prev >>

[Link](https://daily.tinyprojects.dev/182)

---

### 183 —

183 - User reviews | Tiny Projects User reviews Friday, November 11th, 2022 Finally added a feature I've always wanted on PromptBase: user reviews! I've gone for a classic 5-star system with the option for a text review. They look great on the store pages! I also added info about word-count and whether the prompt includes tips. Reviews are a win-win- win for everyone. Sellers get more of a feedback loop from customers. Buyers feel more confident buying a prompt with good reviews. The pressure gets taken off me a lot. I test every prompt to make sure it works, now customers can help me out with weeding out prompts that don't work so well. My only regret is I should have added them sooner! << Next Prev >>

[Link](https://daily.tinyprojects.dev/183)

---

### 185 —

185 - 10,000 Users | Tiny Projects 10,000 Users Tuesday, November 29th, 2022 PromptBase hit a huge milestone: 10k users! I need a nicer way to view this number.. These are registered buyers and sellers on the marketplace. Very exciting!! Growth numbers are a bit skewed right now - after The Verge feature in September there was a large spike in sign ups. Since then, daily signups have actually been declining(!) This is natural though (I could never organically reach the same traffic as that article). Thankfully, it now seems to be back on a normal growth curve again. Still, very pleased with all this considering PromptBase had about 10 users in July. << Next Prev >>

[Link](https://daily.tinyprojects.dev/185)

---

### 186 —

186 - AI Popularity | Tiny Projects AI Popularity Wednesday, November 30th, 2022 It feels like everyday there is some new AI advancement, it's quite hard to keep up! Yesterday a new version of GPT-3 was released. Before that, it was Stable Diffusion 2.0, and the week before that, Midjourney v4 and the DALLE-2 API. I'm in quite a unique position where I can gauge which AI models are popular and producing the best images at any one time based on PromptBase submissions. Hands down, right now, I feel Midjourney v4 is leaps and bounds ahead any other model at producing consistent, usable images from prompts. DALL-E 2 once held this crown, but it's popularity seems to have waned recently, which is kinda sad - it was the OG! Professional Product Photography Prompt by Midjourney v4 << Next Prev >>

[Link](https://daily.tinyprojects.dev/186)

---

### 188 —

188 - December Goals | Tiny Projects December Goals Friday, December 2nd, 2022 Yesterday ended up being very unproductive in the end - the internet in my new place is not good! I ended up buying one of these pre-paid routers to keep me going. It cost £70 with 300GB of data - but seems 10X faster than hotspotting off a phone - it's pretty good. I've been thinking what I want to achieve this month on PromptBase, and decided what I really need to do is start automating away every manual process. Reviewing prompts, social media posts, featuring prompts, approving reviews - these all take too much of my time. I want to try and run as a one-man-army for as long as possible, whilst adding new features to the site (and having a life!). More automation is the way to do this. << Next Prev >>

[Link](https://daily.tinyprojects.dev/188)

---

### 189 —

189 - Max Speed Prompt Reviews | Tiny Projects Max Speed Prompt Reviews Tuesday, December 6th, 2022 Was trying to spot some turtles here 🐢 Added some huge automation improvements to reviewing prompts over the weekend. Now, when a prompt gets submitted, everything I was doing manually is done automatically, like creating thumbnails, and even generating test images using the prompt. Now I can just whiz through, check the test images are good, and click "yes" or "no" if it should be approved for the site. Very excited about this - reviewing prompts was so time intensive, so I now get a couple of hours back to put into other things, and I can onboard users way faster. The final step is making an AI do the final "yes"/"no" approve decision - which I think could be possible! << Next Prev >>

[Link](https://daily.tinyprojects.dev/189)

---

### 190 —

190 - Google Growth | Tiny Projects Google Growth Friday, December 9th, 2022 All of my fancy new prompt review automations have been breaking! Each day I go to review the prompts and something has gone wrong. With each meltdown it slowly improves though I guess.. Something I've been quite excited about the past few days is that Google has been indexing PromptBase pages, and, as such, the SEO is improving quite a bit. 🟣 = impressions 🔵 = clicks Back in June I spent a painful week trying to get Angular Universal working so that the website would be server-side rendered and have better SEO. Battling all those error messages is finally paying off! << Next Prev >>

[Link](https://daily.tinyprojects.dev/190)

---

### 191 —

191 - Super Social | Tiny Projects Super Social Monday, December 12th, 2022 The WiFi at my new place became so slow, that I'm now staying in an AirBnB for a few days whilst better internet gets set up. It's great! I can see and hear the sea from my new tiny desk. This weekend I automated away PromptBase's social media posts. Each day, I still post a single collage of a prompt featured on the site. Some recent Instagram posts However - increasingly, I was finding that I'd get to the end of a day, about to go to sleep, then, realising I forgot to post anything, have to spend 10 tired minutes making a post. Now, at 8 pm each day, my robot social media consultant posts to Instagram and Twitter for me. I just flag prompts that I want it to post. Way easier! It even uses AI to choose an appropriate emoji for the post! I did not choose those emojis! I'm pretty stingy, so I did this through the official APIs instead of paying for a scheduling tool - but, I can see why they exist now. The Facebook/Instagram API in particular is one of the most confusing APIs to set up that I've ever tried. << Next Prev >>

[Link](https://daily.tinyprojects.dev/191)

---

### 192 —

192 - One Year of Paper Website | Tiny Projects One Year of Paper Website Thursday, December 15th, 2022 One year ago I launched Paper Website on Hacker News! Funnily enough, that was also its biggest day of sales, and since then the growth has been flat - well, actually slowly declining... This is mainly my fault for not giving it enough love, and jumping to PromptBase - but hey! That's the whole point of building lots of tiny projects - place tiny bets and stick with the ideas that gain traction. I'm definitely going to keep working on it, but more as a side-side project now. I love using it every day for this daily blog, so it has big value for me. Happy birthday Paper Website! << Next Prev >>

[Link](https://daily.tinyprojects.dev/192)

---

### 193 —

193 - Tiny Press Tour | Tiny Projects Tiny Press tour Monday, December 20th, 2022 Last week PromptBase got quite a lot of press! I did interviews with Fast Company and MadeWithAI , and also had a feature in the Financial Times . I actually own promptbay.com - almost called it that.. I'm under no illusion that this is mainly because selling prompts is somewhat of a strange, taboo topic right now. But hey, all press is good press - and you gotta love a backlink. Some other interesting questions around PromptBase are: - Will the need for prompts and prompt engineers die out as AI gets better? - Are you concerned about "stealing art" used to train models like DALL-E? You can check out those interviews for my answers, but I'll write about them here soon too! << Next Prev >>

[Link](https://daily.tinyprojects.dev/193)

---

### 194 —

194 - Prompt Search 2.0 | Tiny Projects Prompt Search 2.0 Wednesday, December 21st, 2022 Launched a huge update to search on PromptBase yesterday. Vertical scrolling is back! Along with the ability to really niche down on what prompt you want with filters. You can also unravel prompt collections and explore them vertically. And also added better suggestions for similar prompts. Very happy with this change - for a prompt marketplace PromptBase used to have a very poor marketplace page. The site now feels way more fun to navigate, which should hopefully result in more sales for sellers! Already seeing a big uptick in engagement. << Next Prev >>

[Link](https://daily.tinyprojects.dev/194)

---

### 196 —

196 - Prompt Editing | Tiny Projects Prompt Editing Wednesday, January 4th, 2022 Yesterday I launched a feature that allows users to edit their prompts. It sounds insane that you couldn't do this before. When a prompt went live, it was locked. You could only edit it if you sent me an email. I got a lot of emails, and a lot of requests for this feature... The reason it took so long was because this wasn't just a simple database update - there were a few things to consider. Like: How can you ensure an updated prompt works? How can you ensure an update doesn't change the description of the prompt to make it misleading? How do you issue an updated prompt to previous buyers? Basically this was solved by introducing a new "micro review" process if specific updates are made (title, prompt or images) - a bit like how Apple reviews app updates. Some other cool benefits of this feature is that users can edit their own prices now, which should be positive for revenue, e.g. if someone's prompt gets popular they could increase the price, and the opposite if it's not getting many sales. Prompts that get declined can also be edited and re-reviewed, whereas previously you'd have to re-submit a completely new prompt - a lot nicer! Generally I'm finding updates like this require a lot more consideration now. The PromptBase Marketplace is like a living, breathing beast - I don't want to anger it! << Next Prev >>

[Link](https://daily.tinyprojects.dev/196)

---

### 197 —

197 - I broke the site | Tiny Projects I broke the site Monday, January 9th, 2022 Quite badly, for like 12 hours! On Saturday night I was just finishing up a new feature that lets you favourite a prompt. It was getting late, and I had a headache, so I made the mistake of pushing the feature live without much thought. It worked great in my local environment, what could go wrong? I closed my laptop and went to bed. Sunday morning rolls around, I open my inbox to see this: Only half of them 50+ messages about account pages being down. Sellers couldn't see their prompts, buyers couldn't see their purchases. Turns out I'd forgotten to update some database security rules in production for my new "Favourites" feature, and because favourites appear on your account page - it was tanking account pages. It was a 30 second fix. Still quite annoying - because of timezones, those 12 hours I was offline overnight was day time for most people - at least it was the weekend I guess. << Next Prev >>

[Link](https://daily.tinyprojects.dev/197)

---

### 198 —

198 - 994 Button Presses | Tiny Projects 994 Button Presses Wednesday, January 11th, 2022 The feature that broke PromptBase the other day is operational and working well. Now, if you see a prompt you like, you can favorite it by clicking a little heart button. Then your favorite prompts are saved to your account. Pretty basic idea - however since it launched a few days ago that lil' heart button has been pressed 994 times! It's a win-win-win feature for all parties involved: Buyers can browse longer and use the site even if not purchasing. It's also useful for bookmarking prompts to come back to later. Sellers get more feedback on prompts, and favorites serve as social proof to increase sales. For PromptBase, favorites can be used to better rank prompts to highlight the best ones. << Next Prev >>

[Link](https://daily.tinyprojects.dev/198)

---

### 199 —

199 - TwitBase | Tiny Projects TwitBase Friday, January 13th, 2022 PromptBase is a marketplace, but in a very strange way it's also like a social network for Prompt Engineers. There's profiles, messaging, likes - the only thing missing really was followers and the ability to follow other creators. ...Until now! With one click you can follow a creator and get notified whenever they release a new prompt (I also gave profile pages a fresh lick of paint) At the time of writing I have a whopping 12 followers - practically a celebrity. << Next Prev >>

[Link](https://daily.tinyprojects.dev/199)

---

### 200 —

200 - Two Hundred | Tiny Projects Two Hundred Wednesday, January 18th, 2022 Two hundred paper-based blog posts! Thanks so much if you've ever read just one. For me these daily blog posts keep me accountable, motivated, and are great to look back on to measure progress. 100 blog posts ago I was struggling to work out how to grow Paper Website even by a few users. Today PromptBase will hurtle pass 16k users! I've never been more confident in the approach of trying lots of tiny ideas until one shows promise. Thanks for coming along for the ride. Here's to 100 more! << Next Prev >>

[Link](https://daily.tinyprojects.dev/200)

---

### 201 —

201 - Where I've been | Tiny Projects Where I've been Wednesday, October 9th, 2024 It has been way too long! After hitting 200 daily blog posts, and toasting to 100 more, this daily blog just stopped with no notice. I've gotten a few emails asking if I'm okay, so I thought I'd dust off the pen and paper and answer some questions. What are you up to? First off, I am absolutely fine! I have spent the past 18 months pretty much working on PromptBase. The day I posted my last blog update was the day I decided to hand in my notice at my old job and attempt to give this indie hacking thing a go for real. It's been awesome. I spend all day coding; basically living the dream! I feel very lucky to be able to do this. Why did you stop blogging? It was a few reasons: 1. I took a break from blogging whilst switching to full-time on PromptBase, then I lost the routine. 2. I'm a lot more conscious about sharing what I'm working on because of competition. 3. I was focusing on building and, although this blog is fun, it took time away from that. How is PromptBase doing? Right now it's at 249K users and doing well. I have added a lot of new features since my last post: <https://promptbase.com/changelog> . How are your other projects ? Mailoji and Paper Website are still going. They are both in maintenance mode, so I don't really work on them, but they still bring in a few hundred dollars each month. Where are you in the world? My time in Australia ended last Summer, and I'm now back in the UK. Currently I live in a town called Cheltenham - if you're nearby, let's grab a coffee! What are your plans for Tiny Projects in the future? Right now I'm not sure! All I know is I'd like to write a blog post one day about PromptBase (a bit like what I've done with my other projects), but I'll probably do that once: 1. I sell it 2. I get it to a stable point feature-wise. 3. It crashes and burns for some reason. Right now, though, I'm enjoying just focusing on building. ----------- So that's it, really! I'm so sorry for zero updates. I'm not planning on writing daily blogs here for a while, but thanks so much if you're still reading; I massively appreciate it - you're awesome. It was good to catch up! Prev >>

[Link](https://daily.tinyprojects.dev/201)

---
