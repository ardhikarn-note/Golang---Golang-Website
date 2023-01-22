package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("contact page"))
	})
	mux.HandleFunc("/get-post", handler.GetPostHandler)
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/process", handler.Process)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
