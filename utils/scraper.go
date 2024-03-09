package utils

import (
	"fmt"

	"github.com/gocolly/colly"
)

func SetupScrapers(c *colly.Collector, linksFile, textFile, imageDir string) {
	// Create files
	CreateFile(linksFile)
	CreateFile(textFile)

	// Set up callbacks for scraping data
	setupPhoneScraper(c, textFile)
	setupLinkScraper(c, linksFile)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
	})
}

func setupPhoneScraper(c *colly.Collector, textFile string) {
	c.OnHTML("a[href^='tel:']", func(e *colly.HTMLElement) {
		phoneNumber := e.Text
		saveNumber := "Phone No: " + phoneNumber
		fmt.Println("Phone Number:", phoneNumber)
		WriteToFile(textFile, saveNumber)
	})
}

func setupLinkScraper(c *colly.Collector, linksFile string) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link:", link)
		WriteToFile(linksFile, link)
	})

}
