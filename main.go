package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	linksFile, err := os.Create("links.txt")
	if err != nil {
		fmt.Println("Error creating links file:", err)
		return
	}
	defer linksFile.Close()

	textFile, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Error creating text file:", err)
		return
	}
	defer textFile.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("www.wikipedia.org", "example.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		text := e.Text
		_, err := linksFile.WriteString(fmt.Sprintf("Link found: %s\nText: %s\n\n", link, text))
		if err != nil {
			fmt.Println("Error writing to links file:", err)
		}

		e.Request.Visit(link)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := e.Text
		_, err := textFile.WriteString(fmt.Sprintf("Text from <p> tag: %s\n\n", text))
		if err != nil {
			fmt.Println("Error writing to text file:", err)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
	})

	c.Visit("https://www.wikipedia.org")
	c.Visit("https://example.com")
}
