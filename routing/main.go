package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	}

	// http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	// using anonymous function
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Again"))
	})

	// routing using static assets
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.Handle("/", http.FileServer(http.Dir("assets")))
	http.Handle("/static", http.FileServer(http.Dir("assets")))

	// running server
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
