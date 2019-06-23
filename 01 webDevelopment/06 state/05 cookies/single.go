package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER!")
}

func readCookie(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("my-cookie")

	if e != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	fmt.Fprintln(res, "YOUR COOKIE:", cookie)
}
