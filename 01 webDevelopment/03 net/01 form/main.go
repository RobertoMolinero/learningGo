package main

import (
	"html/template"
	"log"
	"net/http"
)

type input string

func (h input) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	e := req.ParseForm()
	if e != nil {
		log.Fatalln(e)
	}

	tpl.Execute(res, req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d input
	http.ListenAndServe(":8080", d)
}
