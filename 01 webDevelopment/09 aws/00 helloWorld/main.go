package main

import (
	"io"
	"log"
	"net/http"
)

type user struct {
	Name string
}

func (u user) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hola "+u.Name+", que haces el fin de semana?")
}

func main() {
	var u = user{
		Name: "Guapa",
	}

	e := http.ListenAndServe(":8080", u)

	if e != nil {
		log.Fatalln(e)
	}
}
