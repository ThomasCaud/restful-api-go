package main

import (
	"github.com/ThomasCaud/go-rest-api/handler"
	"github.com/gorilla/mux"
)

func getRouter(env *Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", env.db.handler.Get).Methods("GET")
	router.HandleFunc("/books", handler.Create).Methods("POST")
	router.HandleFunc("/books/{id}", handler.Delete).Methods("DELETE")
	router.HandleFunc("/books/{id}", handler.Put).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.GetItem).Methods("GET")

	return router
}
