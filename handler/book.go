package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/google/uuid"
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
		{"/books", getCollection(app), "GET"},
		{"/books", create(app), "POST"},
		{"/books/{id}", delete(app), "DELETE"},
		{"/books/{id}", put(app), "PUT"},
		{"/books/{id}", getItem(app), "GET"},
	}
}

func getUuidFromVar(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return id, err
}

func getCollection(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := app.BooksDatabase.GetBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}

func getItem(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getUuidFromVar(w, r)
		if err != nil {
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

func create(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var book model.Book
		json.Unmarshal(reqBody, &book)

		id, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		book.Id = id

		err = app.BooksDatabase.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}
}

func delete(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getUuidFromVar(w, r)
		if err != nil {
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

func put(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var updatedBook model.Book
		json.Unmarshal(reqBody, &updatedBook)

		id, err := getUuidFromVar(w, r)
		if err != nil {
			return
		}
		updatedBook.Id = id

		err = app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(updatedBook)
	}
}
