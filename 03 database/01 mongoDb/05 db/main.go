package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {
	session, e := mgo.Dial("mongodb://localhost")

	if e != nil {
		log.Fatalln(e)
	}

	return session
}
