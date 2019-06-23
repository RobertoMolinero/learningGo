package main

import (
	"io"
	"net/http"
)

type top struct{}

func (top) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "top tip top")
}

type itchy struct{}

func (itchy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "peep peep peep")
}

type scratchy struct{}

func (scratchy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	var t top
	var i itchy
	var s scratchy

	mux := http.NewServeMux()
	mux.Handle("/", t)
	mux.Handle("/itchy", i)
	mux.Handle("/scratchy/", s)

	http.ListenAndServe(":8080", mux)
}
