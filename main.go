package main

import (
	"fmt"
	"golang-web-dev/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Print("\x1bc")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Login)
	mux.HandleFunc("/logout", handler.Logout)
	mux.HandleFunc("/insert", handler.Insert)

	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/auth", handler.Auth)
	mux.HandleFunc("/key", handler.KeyHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/task", handler.Task)
	mux.HandleFunc("/home", handler.HomeHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fileImage := http.FileServer(http.Dir("img"))
	mux.Handle("/public/", http.StripPrefix("/public", fileImage))

	// mux.Handle("/static", http.FileServer(http.Dir("./site")))

	port := func() string {
		p := os.Getenv("PORT")
		if p == "" {
			return "8080"
		}
		return p
	}()

	log.Println("Starting web on port " + port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)

}
