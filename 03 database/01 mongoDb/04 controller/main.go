package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserController()

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", router)
}
