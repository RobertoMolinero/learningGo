package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	glob, e := template.ParseGlob("../templates/*.gohtml")
	handleError(e)

	e = glob.Execute(os.Stdout, nil)
	handleError(e)

	e = glob.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	handleError(e)
}

func handleError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
