package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (i Info) GetAffiliationInfo() string {
	return "have 98 members"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Yulyano Thomas Djaya",
			Gender:  "Male",
			Hobbies: []string{"Gaming", "Coding"},
			Info:    Info{"Lepkom", "Depok"},
		}

		tmpl := template.Must(template.ParseFiles("views/index.html"))

		err := tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Starting server at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
