package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gorilla/mux"
)

type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
}

var Books []model.Book

func GetBooks(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := app.BooksDatabase.GetBooks()
		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, http.StatusText(500), 500)
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
			http.Error(w, "Id must be an integer", 500)
			return
		}

		book, err := app.BooksDatabase.GetBook(id)
		if err != nil {
			http.Error(w, "Not found.", 404)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

/*
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
*/
