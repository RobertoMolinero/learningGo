package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", uploadServeHTTP)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func uploadServeHTTP(res http.ResponseWriter, req *http.Request) {

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	e := tpl.ExecuteTemplate(res, "index.gohtml", body)

	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
		log.Fatalln(e)
	}
}
