package main

import (
	"./controllers"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var root = flag.String("root", ".", "file system path")

func main() {
	router := httprouter.New()
	lc := controllers.NewLineController(getSession(), getTemplate())

	router.GET("/line", lc.GetLines)
	router.POST("/line", lc.CreateLine)
	router.POST("/deleteLine", lc.DeleteLine)
	router.POST("/updateLine", lc.UpdateLine)

	router.ServeFiles("/static/*filepath",http.Dir("static"))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	e := http.ListenAndServe(":8080", router)
	checkErrorWithLogFatalln(e)
}

func getSession() *sql.DB {
	db, e := sql.Open("mysql", "<USER>:<PASSWORD>@tcp(localhost:3306)/goMaria?charset=utf8")
	checkErrorWithLogFatalln(e)

	e = db.Ping()
	checkErrorWithLogFatalln(e)

	return db
}

func getTemplate() *template.Template {
	return template.Must(template.ParseGlob("templates/*"))
}

func checkErrorWithLogFatalln(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
