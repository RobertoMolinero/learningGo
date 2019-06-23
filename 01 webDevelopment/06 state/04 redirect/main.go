package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", fooServeHTTP)
	http.HandleFunc("/bar", barServeHTTP)
	http.HandleFunc("/barred", barredServeHTTP)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func fooServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func barServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method, "\n\n")

	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusSeeOther)
}

func barredServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method, "\n\n")
	e := tpl.ExecuteTemplate(res, "index.gohtml", nil)

	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
		log.Fatalln(e)
	}
}
