package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	HOST := "localhost:8000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.html")

		tmpl, err := template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := map[string]interface{}{
			"title": "Learning Go web",
			"name":  "Yulyano Thomas Djaya",
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Starting server at " + HOST)
	http.ListenAndServe(HOST, nil)

}
