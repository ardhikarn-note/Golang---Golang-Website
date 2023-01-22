package handler

import (
	"golangweb/entity"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{ID: 1, Name: "Product 1", Price: 10000, Stock: 11},
		{ID: 2, Name: "Product 2", Price: 20000, Stock: 3},
		{ID: 3, Name: "Product 3", Price: 30000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Home Page",
		"content": "This is home page",
		"views":   9000,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("GET method"))
	case "POST":
		w.Write([]byte("POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("msg")

		w.Write([]byte("Name: " + name + ", Message: " + message))
		return
	}
}
