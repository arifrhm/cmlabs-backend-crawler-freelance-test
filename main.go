package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	crawlWebsites()
}

func crawlWebsites() {
	targetURLs := []string{
		"https://cmlabs.co",
		"https://sequence.day",
		// Add more URLs
	}

	for _, url := range targetURLs {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching URL %s: %s\n", url, err.Error())
			continue
		}
		defer response.Body.Close()

		htmlBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading HTML from URL %s: %s\n", url, err.Error())
			continue
		}

		html := string(htmlBytes)
		fileName := generateFileName(url)
		err = saveHTMLToFile(fileName, html)
		if err != nil {
			fmt.Printf("Error saving HTML to file %s: %s\n", fileName, err.Error())
			continue
		}

		fmt.Printf("Crawled and saved %s\n", url)
	}
}

func generateFileName(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		// Handle the error if the URL is invalid
		return "Invalid URL"
	}

	domain := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	return domain + ".html"
}

func saveHTMLToFile(fileName string, html string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(html)
	if err != nil {
		return fmt.Errorf("error saving HTML to file %s: %w", fileName, err)
	}

	return nil
}
