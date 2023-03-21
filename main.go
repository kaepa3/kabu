package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Target URL
	url := "https://finance.yahoo.co.jp/quote/998407.O"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	// Instantiate default collector
	c := colly.NewCollector()

	// Extract title element
	c.OnHTML("div[class=_3DSN3qZQ]", func(h *colly.HTMLElement) {
		h.ForEach("span[class=_3wVTceYe]", func(i int, e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	})

	// Start scraping on https://XXX
	c.Visit(url)
}
