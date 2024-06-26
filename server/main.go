package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
	case "/body":
		body(w, r)
	default:
		fmt.Fprint(w, "Error")
	}
	fmt.Printf("Handling function with %s request \n", r.Method)
}

func anotherPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style=\"background-color:green;\"> Hello World <h1>")
}
func main() {

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxCustomMode http.ServeMux
	server.Handler = &muxCustomMode
	muxCustomMode.HandleFunc("/", anotherPage)
	server.ListenAndServe()
}
