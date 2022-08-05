package main

import (
	"fmt"
	"golang-web-dev/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Print("\x1bc")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/key", handler.KeyHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fileImage := http.FileServer(http.Dir("img"))
	mux.Handle("/public/", http.StripPrefix("/public", fileImage))

	// mux.Handle("/static", http.FileServer(http.Dir("./site")))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
