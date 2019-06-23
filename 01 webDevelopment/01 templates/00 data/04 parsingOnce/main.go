package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {
	e := tpl.Execute(os.Stdout, nil)
	if e != nil {
		log.Fatalln(e)
	}

	e = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}

	e = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}

	e = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}
