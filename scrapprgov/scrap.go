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
			var tittle,status,agency,location,opendate,preauctiondate,sheetdate string
			tittle = e.ChildText("span.title")
			status = e.ChildText("div.agency") //status esta como: class = agency, arreglar con una condicion
			agency = e.ChildText("div.agency")
			location = e.ChildText("div.localization")
			opendate = e.ChildText("div.fechaApertura")
			preauctiondate = e.ChildText("div.fechaPreSubasta")
			sheetdate = e.ChildText("div.fechaPliegos")
			if tittle == ""{
				return
			}
			fmt.Printf("Tender tittle: %s \n",tittle)
			fmt.Printf("%s\n",status)
			fmt.Printf("%s\n",agency)
			fmt.Printf("%s\n",location)
			fmt.Printf("%s\n",opendate)
			fmt.Printf("%s\n",preauctiondate)
			fmt.Printf("%s\n",sheetdate)
		})
	})

	c.Visit("https://www2.pr.gov/subasta/Pages/subastas.aspx")

}