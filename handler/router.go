package handler

import (
	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/gorilla/mux"
)

type App struct {
	BooksDatabase db.BooksDatabaseImpl
}

func GetRouter(app *App) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", GetBooks(app)).Methods("GET")
	/*router.HandleFunc("/books", app.Create).Methods("POST")
	router.HandleFunc("/books/{id}", app.Delete).Methods("DELETE")
	router.HandleFunc("/books/{id}", app.Put).Methods("PUT")*/
	router.HandleFunc("/books/{id}", GetBook(app)).Methods("GET")

	return router
}
