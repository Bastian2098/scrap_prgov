package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type tenderdata struct{
	tittle string
	status string
	agency string
	location string 
	openDate string
	preAuctionDate string
	sheetDate string
}

func main() {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando ", r.URL)
	})

	c.OnHTML("div.cbq-layout-main",func (e *colly.HTMLElement){
		e.ForEach("li.pageItemIndicator.listitem", func (_ int, e *colly.HTMLElement) {
			var tendertittle string
			tendertittle = e.ChildText("span.title")
			if tendertittle == ""{
				return
			}
			fmt.Printf("Tender tittle: %s \n",tendertittle)
		})
	})

	c.Visit("https://www2.pr.gov/subasta/Pages/subastas.aspx")

}