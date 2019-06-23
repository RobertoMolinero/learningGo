package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("session")

	if e != nil {
		uuid, _ := uuid.NewV4()

		cookie = &http.Cookie{
			Name:     "session",
			Value:    uuid.String(),
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie)
}
