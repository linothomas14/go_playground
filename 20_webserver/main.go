package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	msg := "Index"

	w.Write([]byte(msg))
}

func handlerSayHello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello"
	w.Write([]byte(msg))
}

func main() {

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerSayHello)

	host := "localhost:8000"
	fmt.Println("Run at : http://" + host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
