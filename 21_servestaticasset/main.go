package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	host := "localhost:8000"
	fmt.Println("Run at : http://" + host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
