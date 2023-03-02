package scrapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type dataStruct struct {
	Name      string
	Position  string
	Office    string
	Age       string
	StartDate string
	Salary    string
}

var employees []dataStruct

func ScrapDataTables() {

	c := colly.NewCollector()

	c.OnHTML("table#example > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			data := dataStruct{
				Name:      el.ChildText("td:nth-child(1)"),
				Position:  el.ChildText("td:nth-child(2)"),
				Office:    el.ChildText("td:nth-child(3)"),
				Age:       el.ChildText("td:nth-child(4)"),
				StartDate: el.ChildText("td:nth-child(5)"),
				Salary:    el.ChildText("td:nth-child(6)"),
			}

			employees = append(employees, data)
			// fmt.Println(el.ChildText("td:nth-child(1)"))
		})

	})

	c.Visit("https://datatables.net/examples/styling/display.html")

	data, err := json.MarshalIndent(employees, "", " ")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile("employees.json", data, 0644)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Total data : %d ", len(employees))
}
