package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    string `json:"Id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

var Books []Book

func populateBooks() {
	Books = []Book{
		{Id: "1", Title: "Cracking the coding interview", Price: 40},
		{Id: "2", Title: "Never split the difference", Price: 30},
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	found := false
	for _, book := range Books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
			found = true
		}
	}

	if !found {
		http.NotFound(w, r)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// todo when db: auto generate id
	reqBody, _ := ioutil.ReadAll(r.Body)

	var book Book
	json.Unmarshal(reqBody, &book)

	Books = append(Books, book)

	json.NewEncoder(w).Encode(Books)
	// todo return 201
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
			handleNoContent(w)
		}
	}

	http.NotFound(w, r)
}

func put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newBook Book
	json.Unmarshal(reqBody, &newBook)

	found := false
	for index, book := range Books {
		if book.Id == id {
			Books[index].Title = newBook.Title
			Books[index].Price = newBook.Price
			found = true
		}
	}

	if !found {
		http.NotFound(w, r)
	}
}
