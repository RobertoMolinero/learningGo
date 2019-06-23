package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	userName     string
	lastActivity time.Time
}

var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	user := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", user)
}

func bar(res http.ResponseWriter, req *http.Request) {
	user := getUser(res, req)

	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if user.Role != "007" {
		http.Error(res, "You must be 007 to enter the bar.", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "bar.gohtml", user)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		password := req.FormValue("password")
		firstName := req.FormValue("firstname")
		lastName := req.FormValue("lastname")
		role := req.FormValue("role")

		if _, ok := dbUsers[userName]; ok {
			http.Error(res, "Username already taken.", http.StatusForbidden)
			return
		}

		sessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(res, cookie)

		dbSessions[cookie.Value] = session{userName, time.Now()}

		encryptedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if e != nil {
			http.Error(res, "Password Encryption Error.", http.StatusInternalServerError)
			return
		}

		u := user{
			UserName: userName,
			Password: encryptedPassword,
			First:    firstName,
			Last:     lastName,
			Role:     role,
		}

		dbUsers[userName] = u

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		password := req.FormValue("password")

		u, ok := dbUsers[userName]

		if !ok {
			http.Error(res, "Username and/or password do not match.", http.StatusForbidden)
			return
		}

		e := bcrypt.CompareHashAndPassword(u.Password, []byte(password))

		if e != nil {
			http.Error(res, "Username and/or password do not match.", http.StatusForbidden)
			return
		}

		sessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = session{userName, time.Now()}

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")
	delete(dbSessions, cookie.Value)

	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(res, cookie)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
