package scrapping

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type mhs struct {
	NPM  string
	Nama string
}

var items []mhs

func ScrapBAAK() {

	c := colly.NewCollector()

	c.OnHTML("table.large-only", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			item := mhs{
				NPM:  el.ChildText("td:nth-child(4)"),
				Nama: el.ChildText("td:nth-child(3)"),
			}

			items = append(items, item)
		})
	})
	for kelas := 1; kelas <= 9; kelas++ {
		for page := 1; page <= 4; page++ {
			c.Visit(fmt.Sprintf("https://baak.gunadarma.ac.id/cariMhsBaru?teks=1ia0%d&tipeMhsBaru=Kelas&page=%d", kelas, page))
		}
	}

	data := validateStructNul(items)

	structToJSONFile(data)

	fmt.Println("Total data : ", len(data))
}

func structToJSONFile(data []mhs) {
	file, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("mhs.json", file, 0664)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func validateStructNul(data []mhs) []mhs {

	var result []mhs

	for _, s := range data {
		if s.Nama != "" {
			result = append(result, s)
		}
	}
	return result
}
