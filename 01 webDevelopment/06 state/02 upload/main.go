package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", uploadServeHTTP)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func uploadServeHTTP(res http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)

	if req.Method == http.MethodPost {

		file, header, e := req.FormFile("q")

		if e != nil {
			http.Error(res, e.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		fmt.Println("\nfile:", file, "\nheader:", header, "\nerror", e)

		bytes, e := ioutil.ReadAll(file)

		if e != nil {
			http.Error(res, e.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bytes)

		dst, e := os.Create(filepath.Join("./files/", header.Filename))

		if e != nil {
			http.Error(res, e.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()

		_, e = dst.Write(bytes)

		if e != nil {
			http.Error(res, e.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
