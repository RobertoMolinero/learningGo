package controllers

import (
	"./models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	objectId := bson.ObjectIdHex(id)

	user := models.User{}

	if e := uc.session.DB("go-web-dev-db").C("users").FindId(objectId).One(&user); e != nil {
		res.WriteHeader(http.StatusNotFound)
		return
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

	user.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("users").Insert(user)

	bytes, e := json.Marshal(user)

	if e != nil {
		log.Fatalln(e)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", bytes)
}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	objectId := bson.ObjectIdHex(id)

	if e := uc.session.DB("go-web-dev-db").C("users").RemoveId(objectId); e != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Deleted user", objectId, "\n")
}
