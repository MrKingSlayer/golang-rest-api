package main

import (
	"fmt"
	"net/http"

	"github.com/MrKingSlayer/golang-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)



func main()  {
	router := httprouter.New()
	c := controllers.NewUserController(getSession())
	router.GET("/user/:id", c.GetUser)
    router.POST("/user", c.CreateUser)
    router.DELETE("/user/:id", c.DeleteUser)

	http.ListenAndServe(":8080", router)
	fmt.Println("app listen on port 8080")
}


func getSession() *mgo.Session{
	session, err := mgo.Dial("MONGO_URI")
	if err != nil {
		panic(err)
	}

	return session
}