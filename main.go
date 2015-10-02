package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	tokens  *mgo.Collection
	db      *mgo.Database
)

func main() {
	router := NewRouter()

	// connect mongoDB
	log.Println("Starting mongodb session")
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// instantiate db & tokens collection as variable
	db = session.DB("confessions")
	tokens = session.DB("confessions").C("tokens")

	//cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5000"},
	})

	// start server
	log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
	log.Println("Listening @ localhost:8000")
}
