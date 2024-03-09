package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	file, err := os.Create("links.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("www.wikipedia.org", "example.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		text := e.Text
		_, err := file.WriteString(fmt.Sprintf("Link found: %s\nText: %s\n\n", link, text))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

		e.Request.Visit(link)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
	})

	c.Visit("https://www.wikipedia.org")
	c.Visit("https://example.com")
}
