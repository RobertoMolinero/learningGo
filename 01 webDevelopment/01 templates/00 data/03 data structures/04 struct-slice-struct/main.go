package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/structSliceStruct.gohtml"))
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

	roomster := car{
		Manufacturer: "Skoda",
		Model:        "Roomster",
		Doors:        4,
	}

	mx5 := car{
		Manufacturer: "Mazda",
		Model:        "MX-5",
		Doors:        2,
	}

	sages := []sage{buddha, ghandi, jesus}
	cars := []car{roomster, mx5}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	e := tpl.Execute(os.Stdout, data)

	if e != nil {
		log.Fatalln(e)
	}
}
