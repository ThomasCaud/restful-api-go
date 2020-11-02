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

	bookOutput struct {
		UUID  uuid.UUID `json:"id"`
		Title string    `json:"title"`
		Price int       `json:"price"`
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

// nous sommes dans le package "handler"
func getCollection(app *App) func(c *gin.Context) ([]bookOutput, error) {
	return func(c *gin.Context) ([]bookOutput, error) {

		// une fonction qui appartient au package "service"
		books, err := service.GetCollection(app.BooksDatabase)
		if err != nil {
			return nil, errors.New("Error while trying to get book collection")
		}

		booksOuput := []bookOutput{}
		for _, book := range books {
			booksOuput = append(booksOuput, bookOutput{book.ID, book.Title, book.Price})
		}

		return booksOuput, nil
	}
}

func getItem(app *App) func(c *gin.Context, in *uuidInput) (*bookOutput, error) {
	return func(c *gin.Context, in *uuidInput) (*bookOutput, error) {
		book, err := app.BooksDatabase.GetBook(in.ID)

		if err != nil {
			return nil, errors.NewNotFound(nil, "Book not found.")
		}

		return &bookOutput{book.ID, book.Title, book.Price}, nil
	}
}

type bookPostInput struct {
	titleAndPriceInput
}

func create(app *App) func(c *gin.Context, in *bookPostInput) (*bookOutput, error) {
	return func(c *gin.Context, in *bookPostInput) (*bookOutput, error) {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}
		book := model.Book{ID: id, Title: in.Title, Price: in.Price}

		err = app.BooksDatabase.CreateBook(book)
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}

		return &bookOutput{book.ID, book.Title, book.Price}, nil
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

func put(app *App) func(c *gin.Context, in *bookPutInput) (*bookOutput, error) {
	return func(c *gin.Context, in *bookPutInput) (*bookOutput, error) {
		updatedBook := model.Book{ID: uuid.MustParse(in.ID), Title: in.Title, Price: in.Price}

		err := app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			return nil, errors.New("Error while updating book")
		}

		return &bookOutput{updatedBook.ID, updatedBook.Title, updatedBook.Price}, nil
	}
}
