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
		tenders := make([]tenderdata,0)
		e.ForEach("li.pageItemIndicator.listitem", func (_ int, e *colly.HTMLElement) {	
				tender := new(tenderdata)
				tender.tittle = e.ChildText("span.title")
				tender.status = e.ChildText("div.agency")
			 	tender.agency = e.ChildText("div.agency")
				tender.location = e.ChildText("div.localization")
				tender.openDate = e.ChildText("div.fechaApertura")
				tender.preAuctionDate = e.ChildText("div.fechaPreSubasta")
				tender.sheetDate = e.ChildText("div.fechaPliegos")
				fmt.Printf("Titulo de licitacion: %s \n",tender.tittle)
				fmt.Printf("%s\n",tender.status)
				fmt.Printf("%s\n",tender.agency)
				fmt.Printf("%s\n",tender.location)
				fmt.Printf("%s\n",tender.openDate)
				fmt.Printf("%s\n",tender.preAuctionDate)
				fmt.Printf("%s\n",tender.sheetDate,)
				fmt.Println("")
				tenders = append(tenders,*tender)			
		})
		if tenders == nil{
			fmt.Println("Vacio")
		}else{
			fmt.Println("Numero de licitaciones: ",len(tenders))
		}
	})

	c.Visit("https://www2.pr.gov/subasta/Pages/subastas.aspx")

}