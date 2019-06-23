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
	tpl = template.Must(template.ParseFiles("../templates/sliceStruct.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	ghandi := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	sages := []sage{buddha, ghandi, jesus}
	e := tpl.Execute(os.Stdout, sages)

	if e != nil {
		log.Fatalln(e)
	}
}
