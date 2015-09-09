package main

import (
	"encoding/json"
	"net/http"
    "gopkg.in/mgo.v2/bson"
    "io/ioutil"
    "io"
	"math/rand"
	"fmt"
	"time"

	"github.com/gorilla/mux"
)

func CreateToken() Token {
	// generate new token
	id := shortId(4)
	token := Token{
		id,
		time.Now(),
	}

	// ensure token id doesn't exist
	err := tokens.Find(bson.M{"_id": id}).One(&token)
	if err == nil {
		token = CreateToken()
		return token
	} else {
		return token
	}
}

func GetConfessionToken(w http.ResponseWriter, r *http.Request) {

	res := CreateToken()

	// check if token exists and create a new one
	err := tokens.Find(bson.M{"_id": res.Id}).One(&res)
	if err == nil {
		res.Id = shortId(4)
		fmt.Println(res.Id)
	}

	// save token
	err = tokens.Insert(&res);
	if err != nil {
		panic(err)
	}

	// respond
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func Confess(w http.ResponseWriter, r *http.Request) {
	vars          := mux.Vars(r)
	token         := vars["token"]
	tokenRes      := TokenRes{}
	confession    := Confession{}
	confessionRes := ConfessionRes{}
	confessions   := db.C("confessions")

	err := tokens.Find(bson.M{"_id": token}).One(&tokenRes)

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
	    confessions.Find(nil).Skip(randn).One(&confessionRes)

	    // remove access token and response confession
	    confessions.RemoveId(confessionRes.Id)
	    tokens.RemoveId(tokenRes.Id)

	    confessions.Insert(confession)

	    // send response
	    w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(confessionRes)
	}

}

type Token struct {
	Id      string `bson:"_id", json:"token"`
	Created time.Time
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