package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", setCounterCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setCounterCookie(res http.ResponseWriter, req *http.Request) {

	cookie, e := req.Cookie("counter-cookie")

	if e == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter-cookie",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count = count + 1
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)
}
