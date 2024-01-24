package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type vars struct {
	Name string
	pwd  string
}

func main() {
	home := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	submitForm := func(w http.ResponseWriter, r *http.Request) {
		data := vars{
			Name: r.PostFormValue("name"),
			pwd:  r.PostFormValue("pwd"),
		}
		fmt.Println(data.Name)
		fmt.Println(data.pwd)

		//html to insert table
		templ, _ := template.New("login").Parse("hello {{ .Name }}")
		templ.Execute(w, data)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submitForm)

	http.ListenAndServe(":8080", nil)
}
