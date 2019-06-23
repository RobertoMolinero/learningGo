package main

import (
	"io"
	"net/http"
)

func rootLevelServeHTTP(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		io.WriteString(res, "GET")
	}

	io.WriteString(res, "top tip top")
}

func nodeXServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "peep peep peep")
}

func nodeYServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	http.HandleFunc("/", rootLevelServeHTTP)
	http.HandleFunc("/itchy", nodeXServeHTTP)
	http.HandleFunc("/scratchy/", nodeYServeHTTP)

	http.ListenAndServe(":8080", nil)
}
