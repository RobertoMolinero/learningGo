package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
	cookie, e := req.Cookie("session")

	if e != nil {
		sessionID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
	}

	cookie.MaxAge = sessionLength
	http.SetCookie(res, cookie)

	var u user

	if userName, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[userName.userName]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, e := req.Cookie("session")

	if e != nil {
		return false
	}

	userName := dbSessions[cookie.Value]
	_, ok := dbUsers[userName.userName]
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN!")
	showSessions()

	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}

	fmt.Println("AFTER CLEAN!")
	showSessions()
}

func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.userName)
	}
	fmt.Println("********")
}
