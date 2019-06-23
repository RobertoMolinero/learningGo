package main

import (
	"html/template"
	"log"
	"os"
)

var tplComparison *template.Template

func init() {
	tplComparison = template.Must(template.ParseFiles("../templates/tpl_comparison.gohtml"))
}

func main() {
	score := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	e := tplComparison.Execute(os.Stdout, score)

	if e != nil {
		log.Fatalln(e)
	}
}
