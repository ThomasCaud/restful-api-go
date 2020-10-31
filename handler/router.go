package handler

import (
	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
	"github.com/loopfz/gadgeto/tonic/utils/swag"

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

	tonic.SetErrorHook(jujerr.ErrHook)
	booksHandlers := GetBooksHandlers(app)

	for _, handler := range booksHandlers {
		router.Handle(handler.method, handler.path, tonic.Handler(handler.f, handler.status))
	}
	router.GET("/swagger.json", swag.Swagger(router, "Books API", swag.Version("v1.0")))

	return router
}
