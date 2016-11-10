package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/scvillon/workshop-uva/handlers"
	"github.com/scvillon/workshop-uva/models"

	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Problema ao se conectar com o Mongo!")
		panic(err)
	}
	defer session.Close()
	db := session.DB("test").C("people")

	err = db.DropCollection()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Insert(
		&models.Person{Name: "Ale", Phone: "+55 53 8116 9639"},
		&models.Person{Name: "Cla", Phone: "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	// Register handlers

	// inicia o muxer
	m := pat.New()
	h := handlers.Handlers{Db: db}
	// obs: higher order functions =D
	m.Get("/find/:name", http.HandlerFunc(h.FindPersonHandler))
	m.Get("/people", http.HandlerFunc(h.ListPeopleHandler))
	m.Post("/create", http.HandlerFunc(h.CreatePersonHandler))
	http.Handle("/", m)
	http.ListenAndServe(":8080", nil)
}
