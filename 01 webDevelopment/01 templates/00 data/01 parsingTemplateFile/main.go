package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	files, e := template.ParseFiles("../templates/template.gohtml")

	if e != nil {
		log.Fatalln("Error parsing File", e)
	}

	e = files.Execute(os.Stdout, nil)

	if e != nil {
		log.Fatalln("Error parsing File", e)
	}
}
