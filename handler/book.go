package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gorilla/mux"
)

type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
	DeleteBook() (model.Book, error)
}

var Books []model.Book

func GetBooks(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := app.BooksDatabase.GetBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}

func GetBook(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Id must be an integer.", http.StatusBadRequest)
			return
		}

		book, err := app.BooksDatabase.GetBook(id)
		if err != nil {
			http.Error(w, "Not found.", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

func CreateBook(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo when db: auto generate id
		reqBody, _ := ioutil.ReadAll(r.Body)

		var book model.Book
		json.Unmarshal(reqBody, &book)

		err := app.BooksDatabase.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}
}

func DeleteBook(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Id must be an integer.", http.StatusBadRequest)
			return
		}

		err = app.BooksDatabase.DeleteBook(id)
		if err != nil {
			http.Error(w, "Not found.", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

/*
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
*/
