package handler

import (
	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/errors"
)

type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
	DeleteBook() (model.Book, error)
}

var Books []model.Book

type (
	UuidInput struct {
		Id string `path:"id" validate:"required" description:"UUID"`
	}

	BookPostInput struct {
		Title string `json:"title" validate:"required"`
		Price int    `json:"price" validate:"required,gt=0"`
	}

	BookPutInput struct {
		UuidInput
		Title string `json:"title" validate:"required"`
		Price int    `json:"price" validate:"required,gt=0"`
	}

	BookOutput struct {
		Uuid  uuid.UUID `json:"id"`
		Title string    `json:"title"`
		Price int       `json:"price"`
	}
)

func GetBooksHandlers(app *App) []Handler {
	return []Handler{
		{"/books", getCollection(app), "GET", 200},
		{"/books", create(app), "POST", 201},
		{"/books/:id", delete(app), "DELETE", 204},
		{"/books/:id", put(app), "PUT", 200},
		{"/books/:id", getItem(app), "GET", 200},
	}
}

func getCollection(app *App) func(c *gin.Context) ([]BookOutput, error) {
	return func(c *gin.Context) ([]BookOutput, error) {
		books, err := app.BooksDatabase.GetBooks()
		if err != nil {
			return nil, errors.New("Error while getting books.")
		}

		booksOuput := []BookOutput{}
		for _, book := range books {
			booksOuput = append(booksOuput, BookOutput{book.Id, book.Title, book.Price})
		}

		return booksOuput, nil
	}
}

func getItem(app *App) func(c *gin.Context, in *UuidInput) (*BookOutput, error) {
	return func(c *gin.Context, in *UuidInput) (*BookOutput, error) {
		book, err := app.BooksDatabase.GetBook(in.Id)

		if err != nil {
			return nil, errors.NewNotFound(nil, "Book not found.")
		}

		return &BookOutput{book.Id, book.Title, book.Price}, nil
	}
}

func create(app *App) func(c *gin.Context, in *BookPostInput) (*BookOutput, error) {
	return func(c *gin.Context, in *BookPostInput) (*BookOutput, error) {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}
		book := model.Book{Id: id, Title: in.Title, Price: in.Price}

		err = app.BooksDatabase.CreateBook(book)
		if err != nil {
			return nil, errors.Errorf("Error while generating UUID.")
		}

		return &BookOutput{book.Id, book.Title, book.Price}, nil
	}
}

func delete(app *App) func(c *gin.Context, in *UuidInput) error {
	return func(c *gin.Context, in *UuidInput) error {
		_, err := app.BooksDatabase.GetBook(in.Id)
		if err != nil {
			return errors.NewNotFound(nil, "Book not found")
		}

		err = app.BooksDatabase.DeleteBook(in.Id)
		if err != nil {
			return errors.New("Error while deleting book.")
		}
		return nil
	}
}

func put(app *App) func(c *gin.Context, in *BookPutInput) (*BookOutput, error) {
	return func(c *gin.Context, in *BookPutInput) (*BookOutput, error) {
		updatedBook := model.Book{Id: uuid.MustParse(in.Id), Title: in.Title, Price: in.Price}

		err := app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			return nil, errors.New("Error while updating book.")
		}

		return &BookOutput{updatedBook.Id, updatedBook.Title, updatedBook.Price}, nil
	}
}
