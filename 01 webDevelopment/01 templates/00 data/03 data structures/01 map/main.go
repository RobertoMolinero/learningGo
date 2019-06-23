package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/map.gohtml"))
}

func main() {
	sages := map[string]string{
		"1": "Tick",
		"2": "Trick",
		"3": "Track",
		"4": "Donald",
		"5": "Dagobert",
	}

	e := tpl.Execute(os.Stdout, sages)
	if e != nil {
		log.Fatalln(e)
	}
}
