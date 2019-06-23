package main

import (
	"html/template"
	"log"
	"os"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/struct.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	e := tpl.Execute(os.Stdout, buddha)

	if e != nil {
		log.Fatalln(e)
	}
}
