package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}
	htmlFilePath := filepath.Join(fmt.Sprintf("static%s.html", r.URL.Path))

	if _, err := os.Stat(htmlFilePath); os.IsNotExist(err) {
		htmlFilePath = "static/index.html"
	}
	tmpl, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		log.Println("file parse:", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Templ.execute:", err)
	}
}
