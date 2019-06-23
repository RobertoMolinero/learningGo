package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("session")

	if e != nil {
		sessionId, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionId.String(),
		}
		http.SetCookie(res, cookie)
	}

	var u user

	if userName, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[userName]
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		firstName := req.FormValue("firstname")
		lastName := req.FormValue("lastname")

		u = user{
			UserName: userName,
			First:    firstName,
			Last:     lastName,
		}

		dbSessions[cookie.Value] = userName
		dbUsers[userName] = u
	}

	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	cookie, e := req.Cookie("session")

	if e != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	userName, ok := dbSessions[cookie.Value]

	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[userName]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
