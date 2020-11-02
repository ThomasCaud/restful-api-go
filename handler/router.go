package handler

import (
	"github.com/ThomasCaud/go-rest-api/store/postgres"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
	"github.com/loopfz/gadgeto/tonic/utils/swag"

	"github.com/gin-gonic/gin"
)

// App contain data and functions to make the app works
type App struct {
	BooksDatabase postgres.BooksDatabaseImpl
}

// Handler is the expected struct of an HTTP call handler
type Handler struct {
	path   string
	f      interface{}
	method string
	status int
}

// GetRouter return gin router for the full API
func GetRouter(app *App) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger.json", swag.Swagger(router, "Books API", swag.Version("v1.0")))

	tonic.SetErrorHook(jujerr.ErrHook)
	booksHandlers := GetBooksHandlers(app)

	for _, handler := range booksHandlers {
		router.Handle(handler.method, handler.path, tonic.Handler(handler.f, handler.status))
	}

	return router
}
