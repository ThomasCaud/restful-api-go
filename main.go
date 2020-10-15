package main

import (
	"log"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/handler"
	"github.com/ThomasCaud/go-rest-api/model"
)

type Env struct {
	db model.Datastore
}

func handleRequests(env *Env) {
	log.Fatal(http.ListenAndServe(":3000", getRouter(env)))
}

func main() {
	// todo remove it
	handler.PopulateBooks()

	// todo get env var
	// default port?
	db, err := model.NewDB("postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	handleRequests(env)
}
