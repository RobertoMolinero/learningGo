package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {
	e := t.ExecuteTemplate(os.Stdout, "template.gohtml", 42)
	if e != nil {
		log.Fatalln(e)
	}
}
