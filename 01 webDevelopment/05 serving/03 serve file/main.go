package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", glassJoeServeHTTP)
	http.HandleFunc("/glass_joe.jpg", glassJoePicServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func glassJoeServeHTTP(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/glass_joe.jpg">`)
}

func glassJoePicServeHTTP(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "glass_joe.jpg")
}
