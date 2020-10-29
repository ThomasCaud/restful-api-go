package handler

import (
	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/loopfz/gadgeto/tonic"

	"github.com/gin-gonic/gin"
)

type App struct {
	BooksDatabase db.BooksDatabaseImpl
}

type Handler struct {
	path   string
	f      interface{}
	method string
	status int
}

func GetRouter(app *App) *gin.Engine {
	router := gin.Default()

	booksHandlers := GetBooksHandlers(app)
	for _, handler := range booksHandlers {
		router.Handle(handler.method, handler.path, tonic.Handler(handler.f, handler.status))
	}

	return router
}
