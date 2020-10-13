package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/books", get).Methods("GET")
	myRouter.HandleFunc("/books", create).Methods("POST")
	myRouter.HandleFunc("/books/{id}", delete).Methods("DELETE")
	myRouter.HandleFunc("/books/{id}", put).Methods("PUT")
	myRouter.HandleFunc("/books/{id}", getItem).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	populateBooks()
	handleRequests()
}
