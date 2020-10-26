package handler

import (
	"net/http"

	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/gorilla/mux"
)

type App struct {
	BooksDatabase db.BooksDatabaseImpl
}

type Handler struct {
	path   string
	f      func(http.ResponseWriter, *http.Request)
	method string
}

func GetRouter(app *App) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	booksHandlers := GetBooksHandlers(app)
	for _, handler := range booksHandlers {
		router.HandleFunc(handler.path, handler.f).Methods(handler.method)
	}

	return router
}
