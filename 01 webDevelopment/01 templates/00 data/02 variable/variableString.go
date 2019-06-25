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
	e := tpl.ExecuteTemplate(os.Stdout, "variableString.gohtml", `Release self-focus; embrace other-focus.`)
	if e != nil {
		log.Fatalln(e)
	}
}
