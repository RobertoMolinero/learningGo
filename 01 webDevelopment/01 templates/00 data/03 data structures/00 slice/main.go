package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/slice.gohtml"))
}

func main() {
	characters := []string{"Tick", "Trick", "Track", "Donald", "Dagobert"}
	e := tpl.Execute(os.Stdout, characters)

	if e != nil {
		log.Fatalln(e)
	}
}
