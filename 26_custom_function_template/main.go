package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("index.html"))

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var funcMap = template.FuncMap{
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"sum": func(n ...int) int {
		sum := 0
		for _, i := range n {
			sum += i
		}
		return sum
	},
}

func main() {

	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
