package scrapping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrappingWiki() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	// Find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		// fmt.Println(links)
		for _, i := range links {
			fmt.Println(i)
		}
	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}
