/* 	Testing the basic concept. The following script will scrape the
	  titles of all of the threads on the first two pages of Hacker News */

package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://news.ycombinator.com/")

	if err != nil {
		panic(err)
	}

	i := 1

	bow.Dom().Find("a.storylink").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(i, s.Text())
		i++
	})

	err = bow.Click("a.morelink")

	if err != nil {
		panic(err)
	}

	j := 31

	bow.Dom().Find("a.storylink").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(j, s.Text())
		j++
	})
	
	// Another neat script. This one does the same thing (for page 1) except it'll visit the links on the front page
	// in descending order from top to bottom and retrieve the h1 header from the link it visits
	
	// bow.Dom().Find("a.storylink").Each(func(_ int, s *goquery.Selection) {
	// 	fmt.Println(i, s.Text())
	//	fmt.Println(s.Attr("href")) // gets you the url
	// 	i++
	// 	bow.Click("a.storylink")
	// 	fmt.Println(bow.Dom().Find("h1").Text())
	// 	bow.Back()
	// })
}
