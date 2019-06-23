package controllers

import (
	"./models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	bytes, e := json.Marshal(user)

	if e != nil {
		log.Fatalln(e)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", bytes)
}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{}

	json.NewDecoder(req.Body).Decode(&user)

	user.Id = "007"

	bytes, e := json.Marshal(user)

	if e != nil {
		log.Fatalln(e)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", bytes)
}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Write code to delete user\n")
}
