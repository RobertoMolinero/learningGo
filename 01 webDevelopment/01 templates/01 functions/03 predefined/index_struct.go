package main

import (
	"html/template"
	"log"
	"os"
)

var tplIndexStruct *template.Template

func init() {
	tplIndexStruct = template.Must(template.ParseFiles("../templates/tpl_index_struct.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Roberto",
	}

	e := tplIndexStruct.Execute(os.Stdout, data)

	if e != nil {
		log.Fatalln(e)
	}
}
