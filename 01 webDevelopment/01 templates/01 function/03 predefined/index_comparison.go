package main

import (
	"html/template"
	"log"
	"os"
)

var tplComparison *template.Template

func init() {
	tplComparison = template.Must(template.ParseFiles("tpl_comparison.gohtml"))
}

func main() {
	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	e := tplComparison.Execute(os.Stdout, g1)

	if e != nil {
		log.Fatalln(e)
	}
}
