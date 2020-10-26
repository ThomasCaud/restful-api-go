package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gorilla/mux"
)

type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
	DeleteBook() (model.Book, error)
}

var Books []model.Book

func GetBooksHandlers(app *App) []Handler {
	return []Handler{
		{"/books", GetBooks(app), "GET"},
		{"/books", CreateBook(app), "POST"},
		{"/books/{id}", DeleteBook(app), "DELETE"},
		{"/books/{id}", PutBook(app), "PUT"},
		{"/books/{id}", GetBook(app), "GET"},
	}
}

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
		book, err := app.BooksDatabase.GetBook(vars["id"])

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

		err := app.BooksDatabase.DeleteBook(vars["id"])
		if err != nil {
			http.Error(w, "Not found.", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func PutBook(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var updatedBook model.Book
		json.Unmarshal(reqBody, &updatedBook)

		vars := mux.Vars(r)
		updatedBook.Id = vars["id"]

		err := app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(updatedBook)
	}
}
