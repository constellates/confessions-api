package main

import (
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

// new stuff start

var (
	session    *mgo.Session
	collection *mgo.Collection
	collection2 *mgo.Collection
)

type Note struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedOn   time.Time     `json:"createdon"`
}

type NoteResource struct {
	Note Note `json:"note"`
}

type NotesResource struct {
	Notes []Note `json:"notes"`
}

// new stuff end

func main() {
	router := NewRouter()

	log.Println("Starting mongodb session")
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("notesdb").C("notes")
	collection2 = session.DB("notesdb").C("confessions")

	log.Fatal(http.ListenAndServe(":8000", router))
}
