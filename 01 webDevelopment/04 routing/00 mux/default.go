package main

import (
	"io"
	"net/http"
)

type topLevel struct{}

func (topLevel) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "top tip top")
}

type nodeA struct{}

func (nodeA) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "peep peep peep")
}

type nodeB struct{}

func (nodeB) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	var t topLevel
	var a nodeA
	var b nodeB

	http.Handle("/", t)
	http.Handle("/itchy", a)
	http.Handle("/scratchy/", b)

	http.ListenAndServe(":8080", nil)
}
