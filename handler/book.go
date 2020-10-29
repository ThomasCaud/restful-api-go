package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BooksDatabase interface {
	GetBooks() ([]model.Book, error)
	GetBook() (model.Book, error)
	DeleteBook() (model.Book, error)
}

var Books []model.Book

type UuidInput struct {
	Id string `path:"id" validate:"required" description:"UUID"`
}

type BookOutput struct {
	Uuid  uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Price int       `json:"price"`
}

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
			return nil, errors.New("Not found.")
		}

		return &BookOutput{book.Id, book.Title, book.Price}, nil
	}
}

func create(app *App) func(c *gin.Context) (*BookOutput, error) {
	return func(c *gin.Context) (*BookOutput, error) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var book model.Book
		json.Unmarshal(reqBody, &book)

		id, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.New("Error while generating random UUID.")
		}
		book.Id = id

		err = app.BooksDatabase.CreateBook(book)
		if err != nil {
			return nil, errors.New("Error while saving the book.")
		}

		return &BookOutput{book.Id, book.Title, book.Price}, nil
	}
}

func delete(app *App) func(c *gin.Context, in *UuidInput) error {
	return func(c *gin.Context, in *UuidInput) error {

		err := app.BooksDatabase.DeleteBook(in.Id)
		if err != nil {
			return errors.New("Not found.")
		}

		return nil
	}
}

func put(app *App) func(c *gin.Context, in *UuidInput) (*BookOutput, error) {
	return func(c *gin.Context, in *UuidInput) (*BookOutput, error) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var updatedBook model.Book
		json.Unmarshal(reqBody, &updatedBook)

		uuid, err := uuid.Parse(in.Id)
		if err != nil {
			return nil, errors.New("Invalid UUID provided.")
		}
		updatedBook.Id = uuid

		err = app.BooksDatabase.PutBook(updatedBook)
		if err != nil {
			return nil, errors.New("Error while updating book.")
		}

		return &BookOutput{updatedBook.Id, updatedBook.Title, updatedBook.Price}, nil
	}
}
