package main

import (
	"fmt"
	"net/http"
)

type hotdog struct{}

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Molinero-Key", "This is ...")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
