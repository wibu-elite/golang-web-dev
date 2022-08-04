package handler

import (
	// "fmt"
	"golang-web-dev/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is Some Problems. Keep Calm", http.StatusInternalServerError)
		return
	}

	// data := entity.Produk{ID: 1, Name: "Rengginang Daisuki Red", Price: 15000, Stock: 100}
	data := []entity.Produk{
		{ID: 1, Name: "Rengginang Gula Aren", Price: 15000, Stock: 45},
		{ID: 2, Name: "Rengginang Trasi", Price: 15000, Stock: 150},
		{ID: 3, Name: "Rengginang Tengiri", Price: 20000, Stock: 3},
	}

	// data := map[string]interface{}{
	// 	"title" : "i'm still learning golang web",
	// 	"content" : "Welcome to My Website",
	// }

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "There is Some Problems. Keep Calm", http.StatusInternalServerError)
		return
	}

	
} 

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, Saya Sedang Belajar Golang Web"))
}

func KeyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Key"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintln(w, "Produk : ", idNumber)
	data := map[string]interface{}{
		"content": idNumber,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "produk.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is Some Problems. Keep Calm", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "There is Some Problems. Keep Calm", http.StatusInternalServerError)
		return
	}
}
