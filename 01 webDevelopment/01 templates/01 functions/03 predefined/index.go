package main

import (
	"html/template"
	"log"
	"os"
)

var tplIndex *template.Template

func init() {
	tplIndex = template.Must(template.ParseFiles("../templates/tpl_index.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	e := tplIndex.Execute(os.Stdout, xs)

	if e != nil {
		log.Fatalln(e)
	}
}
