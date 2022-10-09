package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	HOST := "localhost:8000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello guys"))
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		data := M{"name": "Yulyano Thomas Djaya"} // initial key value untuk dirender di HTML

		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html"))

		err := tmpl.ExecuteTemplate(w, "index", data) // akan merender template index

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

		data := M{"name": "Yulyano Thomas Djaya"}

		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html"))

		err := tmpl.ExecuteTemplate(w, "about", data) // akan merender template about

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting server at : http://" + HOST)
	http.ListenAndServe(HOST, nil)
}
