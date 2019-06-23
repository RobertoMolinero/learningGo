package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", fooServeHTTP)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fooServeHTTP(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}
