package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", octopusServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func octopusServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<img src="https://upload.wikimedia.org/wikipedia/commons/5/5d/Blue-ringed_octopus_%28Hapalochlaena_maculosa%29%2C_Parsley_Bay%2C_Sydney%2C_NSW.jpeg" width="500">
	`)
}
