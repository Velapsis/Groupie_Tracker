package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control","no-store, no cache, must-revalidate, proxy-revalidate ")
		http.ServeFile(w, r, "static/html/Menu.html")

	})

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)

	}
}
