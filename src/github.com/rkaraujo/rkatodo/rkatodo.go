package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Content string
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8212", nil)
}

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	todo := &Todo{"Create a website"}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(rw, todo)
}
