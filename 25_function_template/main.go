package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name  string
	Class string
	Power string
}

func (s Superhero) SayHello(greeting string) string {

	res := "Ucapan " + greeting + " dari Superhero " + s.Name
	return res
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	saitama := Superhero{
		Name:  "Hulk",
		Class: "B+",
		Power: "Unlimited strength",
	}

	tmpl := template.Must(template.ParseFiles("views/index.html"))
	err := tmpl.Execute(w, saitama)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/index", indexHandler)
	fmt.Println("Starting server at : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
