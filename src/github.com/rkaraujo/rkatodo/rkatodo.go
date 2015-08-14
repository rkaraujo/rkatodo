package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Content string
	Created time.Time
}

type ListForm struct {
	TodoList []Todo
}

func main() {
	router := httprouter.New()
	router.GET("/", ListHandler)
	http.ListenAndServe(":8212", router)
}

func ListHandler(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	listForm := &ListForm{ []Todo{ Todo{"Create Website", time.Now()}, Todo{"Learn Go", time.Now()} } }
	
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(rw, listForm)
}
