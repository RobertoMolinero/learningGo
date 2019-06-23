package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", glassJoeServeHTTP)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func glassJoeServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/resources/glass_joe.jpg">`)
}
