package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MrKingSlayer/golang-rest-api/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type UserController struct {
	session *mgo.Session
}


func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (c UserController) GetUser(w http.ResponseWriter , r *http.Request , ps httprouter.Params) {
	id:= ps.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid:= bson.ObjectIdHex(id)

	u:= models.User{}

	if err:= c.session.DB("mongo-golang").C("users").FindId(oid).One(&u) ; err != nil {
		w.WriteHeader(404)
		return
	}

	uJson , err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type" , "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w , "%s\n" , uJson)
}
func (c UserController) CreateUser(w http.ResponseWriter , r *http.Request , _ httprouter.Params) {
	u:= models.User{}

	u.Id = bson.NewObjectId()

	json.NewDecoder(r.Body).Decode(&u)
	c.session.DB("mongo-golang").C("users").Insert(u)


	uJson , err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}
	
	w.Header().Set("Content-type" , "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w , "%s\n" , uJson)

}

func (c UserController) DeleteUser(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	id:= ps.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid:= bson.ObjectIdHex(id)

	if err := c.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil{
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// fmt.Fprintf(w , "Delete user" , oid , "\n")
}

