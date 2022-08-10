package handler

import (
	"fmt"
	// "golang-web-dev/entity"
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	// "text/template"
	// "github.com/go-sql-driver/mysql"
	// "golang.org/x/text/message"
)

type Tasks struct {
	id        int
	nama      string
	nama_task string
	deadline  string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "task"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		nama_task := r.FormValue("nama_task")
		deadline := r.FormValue("deadline")
		insForm, err := db.Prepare("INSERT INTO task(nama, nama_task, deadline) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, nama_task, deadline)
		// log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/task", http.StatusSeeOther)
}
func Task(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	if r.Method == "GET" {
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM task ORDER BY id DESC")
		// fmt.Println(selDB)
		if err != nil {
			panic(err.Error())
		}
		// tsk := Tasks{}
		// res := []Tasks{}
		// for selDB.Next() {
		// 	var id int
		// 	var nama, nama_task, deadline string
		// 	err = selDB.Scan(&id, &nama, &nama_task, &deadline)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// 	tsk.id = id
		// 	tsk.nama = nama
		// 	tsk.nama_task = nama_task
		// 	tsk.deadline = deadline
		// 	res = append(res, tsk)
		// }

		var task = Tasks{}

		// iterate over rows
		for selDB.Next() {
			err = selDB.Scan(&task.id, &task.nama, &task.nama_task, &task.deadline)
			fmt.Println(task)
			if err != nil {
				log.Println(err)
				http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			}
		}

		// tmpl.ExecuteTemplate(w, "Index", res)
		// defer db.Close()
		// data := map[string]interface{}{
		// 	"content": res,
		// }
		// fmt.Print(selDB)
		data := map[string]interface{}{
			"content": task,
		}

		fmt.Print(data)
		tmpl, err := template.ParseFiles(path.Join("views", "task.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		// tmpl.Execute(w, data)
		// err = tmpl.Execute(w, res)

		// fmt.Println(res)
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
		// 	return
		// }
		defer db.Close()
		return

	}
	http.Error(w, "There is Something Wrong. Keep Calm", http.StatusBadRequest)
}
func CheckLoginMiddleware(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("login.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	if string(data) == "Berhasil" {
		return
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	log.Println(r.URL.Path)
	if r.URL.Path != "/home" {
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
	// data := []entity.Produk{
	// 	{ID: 1, Name: "Rengginang Gula Aren", Price: 15000, Stock: 45},
	// 	{ID: 2, Name: "Rengginang Trasi", Price: 15000, Stock: 150},
	// 	{ID: 3, Name: "Rengginang Tengiri", Price: 20000, Stock: 3},
	// }

	// data := map[string]interface{}{
	// 	"title" : "i'm still learning golang web",
	// 	"content" : "Welcome to My Website",
	// }

	tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "There is Some Problems. Keep Calm", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	w.Write([]byte("Hello World, Saya Sedang Belajar Golang Web"))
}

func KeyHandler(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	w.Write([]byte("This is Key"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
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
func PostGet(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Something Wrong With You", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	if r.Method == "GET" || r.Method == "POST" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Worng. Keep Calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		return
	}

	// if r.Method == "POST" {
	// 	tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
	// 	if err != nil {
	// 		log.Println(err)
	// 		http.Error(w, "There is Something Worng. Keep Calm", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	err = tmpl.Execute(w, nil)

	// 	if err != nil {
	// 		log.Println(err)
	// 		http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	return
	// }
	http.Error(w, "There is Something Wrong. Keep Calm", http.StatusBadRequest)
}
func Proses(w http.ResponseWriter, r *http.Request) {
	CheckLoginMiddleware(w, r)
	if r.Method == "POST" || r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}
		// w.Write([]byte(name))
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}

	}
	// http.Error(w, "There is Something Wrong. Keep Calm", http.StatusBadRequest)
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "login.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "There is Something Wrong. Keep Calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "There is Something Wrong. Keep Calm", http.StatusBadRequest)
}

func CreateFile(keyword string) {
	file, err := os.Create("login.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	len, err := file.WriteString(keyword)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}

func Logout(w http.ResponseWriter, r *http.Request) {
	err := os.Remove("login.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed deleting file: %s", err)
	}
	CreateFile("kosong")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "admin" && password == "123" {
			CreateFile("Berhasil")

			http.Redirect(w, r, "/home", http.StatusSeeOther)
		} else {

			http.Error(w, "Wah Otaknya ini yang kena", http.StatusBadRequest)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
