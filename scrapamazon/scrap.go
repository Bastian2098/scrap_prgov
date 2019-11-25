package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}

			fmt.Printf("Product Name: %s \n", productName)
		})
	})

	c.Visit("https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1")
}