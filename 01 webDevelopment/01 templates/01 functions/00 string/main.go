package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("../templates/string.gohtml"))
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

	data := struct {
		Wisdom []sage
	}{
		Wisdom: sages,
	}

	e := tpl.ExecuteTemplate(os.Stdout, "string.gohtml", data)

	if e != nil {
		log.Fatalln(e)
	}
}
