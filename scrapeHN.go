package main

// Scrape the titles and corresponding URLs from the front page of Hacker News

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrapeHN() {

	res, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the thread title and urls
	doc.Find("#hnmain tbody tr .itemlist .athing .title .storylink").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title and url
		title := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("Article %d: %s | %s `\n`", i, title, link)
	})
}

func main() {
	scrapeHN()
}
