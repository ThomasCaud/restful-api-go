package handler

import (
	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/ThomasCaud/go-rest-api/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/errors"
)

// BooksDatabase interface contains function to call in order to get data from an "abstract" database
// todo move this interface in service file
type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
	DeleteBook() (model.Book, error)
}

type (
	uuidInput struct {
		ID string `path:"id" validate:"required" description:"UUID"`
	}

	titleAndPriceInput struct {
		Title string `json:"title" validate:"required"`
		Price int    `json:"price" validate:"required,gt=0"`
	}
)

// GetBooksHandlers export books handlers
// todo remove this function
func GetBooksHandlers(app *App) []Handler {
	return []Handler{
		{"/books", getCollection(app), "GET", 200},
		{"/books", create(app), "POST", 201},
		{"/books/:id", delete(app), "DELETE", 204},
		{"/books/:id", put(app), "PUT", 200},
		{"/books/:id", getItem(app), "GET", 200},
	}
}

func getCollection(app *App) func(c *gin.Context) ([]*model.Book, error) {
	return func(c *gin.Context) ([]*model.Book, error) {

		books, err := service.GetCollection(app.BooksDatabase)
		if err != nil {
			return nil, errors.New("Error while trying to get book collection")
		}

		return books, nil
	}
}

func getItem(app *App) func(c *gin.Context, in *uuidInput) (*model.Book, error) {
	return func(c *gin.Context, in *uuidInput) (*model.Book, error) {
		book, err := app.BooksDatabase.GetBook(in.ID)

		if err != nil {
			return nil, errors.NewNotFound(nil, "Book not found.")
		}

		return book, nil
	}
}

type bookPostInput struct {
	titleAndPriceInput
}

func create(app *App) func(c *gin.Context, in *bookPostInput) (*model.Book, error) {
	return func(c *gin.Context, in *bookPostInput) (*model.Book, error) {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}
		book := model.Book{ID: id, Title: in.Title, Price: in.Price}

		err = app.BooksDatabase.CreateBook(book)
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}

		return &book, nil
	}
}

func delete(app *App) func(c *gin.Context, in *uuidInput) error {
	return func(c *gin.Context, in *uuidInput) error {
		_, err := app.BooksDatabase.GetBook(in.ID)
		if err != nil {
			return errors.NewNotFound(nil, "Book not found")
		}

		err = app.BooksDatabase.DeleteBook(in.ID)
		if err != nil {
			return errors.New("Error while deleting book")
		}
		return nil
	}
}

type bookPutInput struct {
	uuidInput
	titleAndPriceInput
}

func put(app *App) func(c *gin.Context, in *bookPutInput) (*model.Book, error) {
	return func(c *gin.Context, in *bookPutInput) (*model.Book, error) {
		updatedBook := model.Book{ID: uuid.MustParse(in.ID), Title: in.Title, Price: in.Price}

		err := app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			return nil, errors.New("Error while updating book")
		}

		return &updatedBook, nil
	}
}
