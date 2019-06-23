package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, e := template.ParseFiles("../templates/one.gohtml")
	if e != nil {
		log.Fatalln(e)
	}

	e = tpl.Execute(os.Stdout, nil)

	tpl, e = template.ParseFiles("../templates/two.gohtml", "../templates/vespa.gohtml")
	if e != nil {
		log.Fatalln(e)
	}

	e = tpl.ExecuteTemplate(os.Stdout, "../templates/two.gohtml", nil)
	if e != nil {
		log.Fatalln(e)
	}
}
