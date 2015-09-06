package main

import (
	"encoding/json"
	"net/http"
    "gopkg.in/mgo.v2/bson"
    "log"
    "io/ioutil"
    "io"
	"math/rand"
	"fmt"

	"github.com/gorilla/mux"
)

func GetConfessionToken(w http.ResponseWriter, r *http.Request) {
	res := Token{uuid()}
	// save token
	err := tokens.Insert(&res);
	if err != nil {
		panic(err)
	} else {
		log.Printf("Saved access token")
	}
	json.NewEncoder(w).Encode(res)
}

func Confess(w http.ResponseWriter, r *http.Request) {
	vars          := mux.Vars(r)
	token         := vars["token"]
	tokenRes      := TokenRes{}
	confession    := Confession{}
	confessionRes := ConfessionRes{}
	confessions   := db.C("confessions")

	err := tokens.Find(bson.M{"token": token}).One(&tokenRes)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(404)
		json.NewEncoder(w).Encode("Token not found")
	} else {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
	        panic(err)
	    }
	    if err := r.Body.Close(); err != nil {
	        panic(err)
	    }
	    if err := json.Unmarshal(body, &confession); err != nil {
	    	panic(err)
	    }
	    count, err := confessions.Find(nil).Count()
	    randn := rand.Intn(count)
	    fmt.Println(count)
	    fmt.Println(randn)
	    confessions.Insert(confession)
	    confessions.Find(nil).Skip(randn).One(&confessionRes)

	    // remove access token and response confession
	    confessions.RemoveId(confessionRes.Id)
	    tokens.RemoveId(tokenRes.Id)

	    // send response
		json.NewEncoder(w).Encode(confessionRes)
	}

}

type Token struct {
	Token string `json:"token"`
}

type TokenRes struct {
	Id    bson.ObjectId `bson:"_id", json:"_id"`
	Token string        `json:"token"`
}

type Confession struct {
	Confession string `json:"confession"`
}

type ConfessionRes struct {
	Id         bson.ObjectId `bson:"_id", json:"_id"`
	Confession string        `json:"confession"`
}