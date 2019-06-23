package main

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
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

	var u user

	if userName, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[userName]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, e := req.Cookie("session")

	if e != nil {
		return false
	}

	userName := dbSessions[cookie.Value]
	_, ok := dbUsers[userName]
	return ok
}
