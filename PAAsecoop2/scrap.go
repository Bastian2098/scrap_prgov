package main

import(
	"fmt"
	"strings"
	//"context"
	//"log"
	//"time"
	"github.com/gocolly/colly"
	"strconv"
	//"github.com/chromedp/chromedp"	
)

type paa struct{ //Estructura que almacena los detalles mostrados en la lista de PAAs
	entity string
	year string
	contactName string
	contactPhone string
	contactEmail string
	publicationDate string
	totalValue [2]string
	smallerAmount [2]string
	minimunAmount [2]string
	version string
	modificationDate string
	state string
	details []detail
}

type detail struct{ //Detalles de la tabla del PAA
	unspsc string
	description string
	startDate string
	presentationDate string
	duration string
	modality string
	source string
	estimatedValue string
	currentEstimated string
	validity string
	requestStatus string
	hiringUnit string
	location string
	responsibleContact [3]string
}

func main(){
	c := colly.NewCollector()
	url := "https://community.secop.gov.co/Public/App/AnnualPurchasingPlanManagementPublic/Index?currentLanguage=es-CO&Page=login&Country=CO&SkinName=CCE"
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando ", r.URL)
	})

	c.OnHTML("table.VortalGrid", func(e *colly.HTMLElement){
		paas := make([]paa,0)
		details := make([]detail,0)
		var url string
		var aux [2]string
		var aux2 [3]string
		for i := 0 ; i<10 ; i++{ //La iteracion es de 10 ya que son los unicos elementos que mostrara sin hacer click en el paginator
			paa := new(paa)
			paa.entity = e.ChildText("span#spnGridAppEntitySpan1_"+strconv.Itoa(i))
			paa.year = e.ChildText("span#spnGridAppYearSpan1_"+strconv.Itoa(i))
			paa.publicationDate = e.ChildText("span#dtbGridAppPublishingDateBox1_"+strconv.Itoa(i)+"_txt")
			aux[0] = e.ChildText("span#cbxGridAppGlobalValueNumBox1_"+strconv.Itoa(i)) 
			aux[1] = e.ChildText("span#spnGridAppGlobalValueCurrencySpan1_"+strconv.Itoa(i))
			paa.totalValue = aux
			paa.version = e.ChildText("span#spnGridAppVersionSpan1_"+strconv.Itoa(i))
			paa.modificationDate = e.ChildText("span#dtbGridAppModifiedDateBox1_"+strconv.Itoa(i)+"_txt")
			paa.state = e.ChildText("span#spnGridAppStateSpan1_"+strconv.Itoa(i))
			e.ForEach("a#lnkGridAppDetailLink_"+strconv.Itoa(i), func (_ int, e *colly.HTMLElement){
				o := colly.NewCollector()
				url = e.Attr("onclick")
				o.OnRequest(func(r *colly.Request) {
					fmt.Println("Visitando ", r.URL)
				})
				o.OnHTML("div#stphStepPlace", func(e *colly.HTMLElement){
					paa.contactName = e.ChildText("span#spnContactName")
					paa.contactPhone = e.ChildText("span#spnContactPhone")
					paa.contactEmail = e.ChildText("span#spnContactEmail")
					aux[0] = e.ChildText("span#cbxBudgetMenorQuantia")
					aux[1] = e.ChildText("span#spnBudgetMenorQuantiaCurrency")
					paa.smallerAmount = aux
					aux[0] = e.ChildText("span#cbxBudgetMinimaQuantia")
					aux[1] = e.ChildText("span#spnBudgetMinimaQuantiaCurrency")
					paa.minimunAmount = aux
					for j := 0 ; j<10 ; j++{ //Similar a la situacion anterior, solo mostrara 10 detalles del PAA por no poder manipular el paginator
						detail := new(detail)
						detail.unspsc = e.ChildText("span#spnGridAcquisitionsCategorySpan2_"+strconv.Itoa(j))
						detail.description = e.ChildText("span#spnGridAcqDescriptionSpan1_"+strconv.Itoa(j)+"_shortMessage")
						detail.startDate = e.ChildText("span#spnGridAcqBeginDateSpan2_"+strconv.Itoa(j))
						detail.presentationDate = e.ChildText("span#spnGridAcqDueDateSpan2_"+strconv.Itoa(j))
						detail.duration = e.ChildText("span#spnGridAcqDurationSpan4_"+strconv.Itoa(j))
						detail.modality = e.ChildText("span#spnGridAcqTypeSpan_"+strconv.Itoa(j))
						detail.source = e.ChildText("span#spnGridAcqBudgetOriginSpan2_"+strconv.Itoa(j))
						detail.estimatedValue = e.ChildText("span#spnGridAcqTotalValueSpan2_"+strconv.Itoa(j))
						detail.currentEstimated = e.ChildText("span#spnGridAcqValueInActualBudgetSpan2_"+strconv.Itoa(j))
						detail.validity = e.ChildText("span#spnGridAcqFutureBudgetSpan2_"+strconv.Itoa(j))
						detail.requestStatus = e.ChildText("span#spnGridAcqFutureBudgetStateSpan2_"+strconv.Itoa(j))
						detail.hiringUnit = e.ChildText("span#spnGridAcqBusinessOperationSpan2_"+strconv.Itoa(j))
						detail.location = e.ChildText("span#spnGridAcqLocationSpan_"+strconv.Itoa(j))
						aux2[0] = e.ChildText("span#spnGridAcqResponsableContactNameSpan_"+strconv.Itoa(j))
						aux2[1] = e.ChildText("span#spnGridAcqResponsableContactPhoneSpan_"+strconv.Itoa(j))
						aux2[2] = e.ChildText("span#spnGridAcqResponsableContactEmailSpan_"+strconv.Itoa(j))
						detail.responsibleContact = aux2
						details = append(details, *detail)
					}
					paa.details = details				
				})
				o.Visit(fix(url))
			})
			fmt.Printf("%s\n",paa.entity)
			fmt.Printf("%s\n",paa.year)
			fmt.Printf("%s\n",paa.publicationDate)
			fmt.Printf("%s\n",paa.totalValue)
			fmt.Printf("%s\n",paa.version)
			fmt.Printf("%s\n",paa.modificationDate)
			fmt.Printf("%s\n",paa.state)
			fmt.Printf("%s\n",paa.contactName)
			fmt.Printf("%s\n",paa.contactPhone)
			fmt.Printf("%s\n",paa.contactEmail)
			fmt.Printf("%s\n",paa.smallerAmount)
			fmt.Printf("%s\n",paa.minimunAmount)
			fmt.Println("")	 		
			paas = append(paas,*paa)					
		}	
	})
	c.Visit(url)
}

func fix(url string) string{ //funcion que arregla la URL del boton "Detalles" de la lista de PAAs
	cin := strings.SplitN(url,"=",2)
	cin2 := strings.Split(cin[1],"+")
	cin3 := make([]string,0)
	for i := 0; i<len(cin2) ; i++ {
		st := strings.Split(cin2[i],"'")
		cin3 = append(cin3,st[1])
	}
	cin4 := make([]string,0)
	cin4 = append(cin4,"https://community.secop.gov.co")
	cin4 = append(cin4, strings.Join(cin3,""))
	cin5 := strings.Join(cin4,"")
	return cin5
}