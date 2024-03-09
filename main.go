package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	imageDir := "images"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.Mkdir(imageDir, 0755)
	}

	c := colly.NewCollector()

	// Find and extract phone number
	c.OnHTML("a[href^='tel:']", func(e *colly.HTMLElement) {
		phoneNumber := e.Text
		fmt.Println("Phone Number:", phoneNumber)
		textFile.WriteString(fmt.Sprintf("Phone No: %s\n", phoneNumber))
	})

	// Find and extract link
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link:", link)
		linksFile.WriteString(fmt.Sprintf("%s\n", link))
	})

	// Find and extract image URL
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		imageURL := e.Attr("src")
		fmt.Println("Image URL:", imageURL)
		// Download the image
		downloadImage(imageURL, imageDir)
	})

	// Visit the URL and scrape the data
	c.Visit("https://search.brave.com/search?q=ccd+nashik&source=web")
}

func downloadImage(url, dir string) {
	fileName := filepath.Base(url)
	filePath := filepath.Join(dir, fileName)
	if strings.HasPrefix(url, "//") {
		url = "https:" + url
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to download image:", err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Failed to save image:", err)
		return
	}

	fmt.Println("Image downloaded successfully:", filePath)
}
