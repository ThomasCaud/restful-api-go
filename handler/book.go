package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gorilla/mux"
)

var Books []model.Book

func PopulateBooks() {
	Books = []model.Book{
		{Id: "1", Title: "Cracking the coding interview", Price: 40},
		{Id: "2", Title: "Never split the difference", Price: 30},
	}
}

func (db *DB) Get(w http.ResponseWriter, r *http.Request) {
	books, err := db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
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

func Create(w http.ResponseWriter, r *http.Request) {
	// todo when db: auto generate id
	reqBody, _ := ioutil.ReadAll(r.Body)

	var book model.Book
	json.Unmarshal(reqBody, &book)

	Books = append(Books, book)

	json.NewEncoder(w).Encode(Books)
	// todo return 201
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
			HandleNoContent(w)
		}
	}

	http.NotFound(w, r)
}

func Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newBook model.Book
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
