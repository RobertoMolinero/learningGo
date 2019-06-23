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
	FirstName    string
	LastName     string
	Subscription bool
}

func main() {
	http.HandleFunc("/", personServeHTTP)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func personServeHTTP(res http.ResponseWriter, req *http.Request) {

	firstName := req.FormValue("first")
	lastName := req.FormValue("last")
	subscription := req.FormValue("subscription") == "on"

	e := tpl.ExecuteTemplate(res, "index.gohtml", person{
		FirstName:    firstName,
		LastName:     lastName,
		Subscription: subscription,
	})

	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
		log.Fatalln(e)
	}
}
