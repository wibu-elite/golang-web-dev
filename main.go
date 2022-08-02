package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/key", KeyHandler)
	mux.HandleFunc("/product", ProductHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)



}
func HomeHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([] byte("Selamat Datang di Website Learn GO Lenguage"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([] byte("Hello World, Saya Sedang Belajar Golang Web"))
}

func KeyHandler(w http.ResponseWriter, r *http.Request){
	w.Write([] byte("This is Key"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query() .Get("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1{
		http.NotFound(w, r)
		return
	}

	// w.Write([] byte("Halaman Produk"))
	fmt.Fprintf(w, "Produk : %d", idNumber)
}