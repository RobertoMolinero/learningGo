package main

import (
	"io"
	"net/http"
	"os"
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
	file, e := os.Open("glass_joe.jpg")
	if e != nil {
		http.Error(res, "File not found.", http.StatusNotFound)
		return
	}

	defer file.Close()

	stat, e := file.Stat()
	if e != nil {
		http.Error(res, "File not found.", http.StatusNotFound)
		return
	}

	http.ServeContent(res, req, file.Name(), stat.ModTime(), file)
}
