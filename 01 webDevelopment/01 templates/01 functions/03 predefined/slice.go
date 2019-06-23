package main

import (
	"html/template"
	"log"
	"os"
)

var tplSlice *template.Template

func init() {
	tplSlice = template.Must(template.ParseFiles("../templates/tpl_slice.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no belief",
		Admin: false,
	}

	u2 := user{
		Name:  "Ghandi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	e := tplSlice.Execute(os.Stdout, users)

	if e != nil {
		log.Fatalln(e)
	}
}
