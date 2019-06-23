package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/expire", expireCookie)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set">set a cookie</a></h1>`)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: "some value",
		Path:  "/",
	})

	fmt.Fprintln(res, `<h1><a href="/read">read</a></h1>`)
}

func readCookie(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("session")

	if e != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(res, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, cookie)
}

func expireCookie(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("session")

	if e != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(res, cookie)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
