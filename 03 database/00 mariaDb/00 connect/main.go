package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var e error

func main() {
	db, e = sql.Open("mysql", "<USER>:<PASSWORD>@tcp(localhost:3306)/goMaria?charset=utf8")
	checkError(e)

	defer db.Close()

	e = db.Ping()
	checkError(e)

	http.HandleFunc("/", indexServeHTTP)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	e := http.ListenAndServe(":8080", nil)
	checkError(e)
}

func indexServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Succesfully completed.")
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
