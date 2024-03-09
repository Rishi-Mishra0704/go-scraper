package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new Colly collector
	c := colly.NewCollector(
		// Visit only the domains example.com
		colly.AllowedDomains("www.wikipedia.org"),
	)

	// Callback function to be executed when visiting each page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found:", link)
		fmt.Println("Text:", e.Text)
		// Visit the link found recursively
		e.Request.Visit(link)
	})

	// Callback function to be executed when the scraping is completed
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
	})

	// Start scraping
	c.Visit("https://www.wikipedia.org")
}
