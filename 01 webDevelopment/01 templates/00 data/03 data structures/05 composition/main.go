package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/composition.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{Number: "001", Name: "Introduction in ...", Units: "3"},
				{Number: "002", Name: "Advanced Topics in ...", Units: "5"},
				{Number: "003", Name: "Master Lectures in ...", Units: "12"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{Number: "004", Name: "Introduction in ...", Units: "3"},
				{Number: "005", Name: "Advanced Topics in ...", Units: "5"},
				{Number: "006", Name: "Master Lectures in ...", Units: "12"},
			},
		},
	}

	e := tpl.Execute(os.Stdout, y)

	if e != nil {
		log.Fatalln(e)
	}
}
