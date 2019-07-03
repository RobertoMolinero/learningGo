package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"time"
)

var fmap = template.FuncMap{
	"formatAsDate": formatAsDate,
}

func formatAsDate(t time.Time) string {
	return t.Format("2006-01-02")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fmap).ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	lc := controllers.NewLineController(tpl)

	router.GET("/line", lc.GetLines)
	router.POST("/createLine", lc.CreateLine)
	router.POST("/deleteLine", lc.DeleteLine)
	router.POST("/updateLine", lc.UpdateLine)

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	e := http.ListenAndServe(":8080", router)
	checkErrorWithLogFatalln(e)
}

func checkErrorWithLogFatalln(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
