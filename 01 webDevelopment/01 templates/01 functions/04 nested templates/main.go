package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	e := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)

	if e != nil {
		log.Fatalln(e)
	}
}
