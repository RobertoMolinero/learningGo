package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl_index.gohtml"))
}

func main() {

	e := tpl.ExecuteTemplate(os.Stdout, "tpl_index.gohtml", time.Now())

	if e != nil {
		log.Fatalln(e)
	}
}
