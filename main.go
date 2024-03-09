package main

import (
	"web-scraper/utils"

	"github.com/gocolly/colly"
)

func main() {
	linksFile := "links.txt"
	textFile := "text.txt"
	imageDir := "images"

	c := colly.NewCollector()

	// Set up scrapers
	utils.SetupScrapers(c, linksFile, textFile, imageDir)

	// Visit the URL and scrape the data
	c.Visit("https://search.brave.com/search?q=ccd+nashik&source=web")
}
