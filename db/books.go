package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ThomasCaud/go-rest-api/model"
	"github.com/google/uuid"
)

type BooksDatabaseImpl struct {
	DB *sql.DB
}

const tableName = "books"

func (this BooksDatabaseImpl) GetBooks() ([]*model.Book, error) {
	rows, err := this.DB.Query("SELECT * FROM " + tableName)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	defer rows.Close()

	books := make([]*model.Book, 0)
	for rows.Next() {
		book := new(model.Book)
		err := rows.Scan(&book.Id, &book.Title, &book.Price)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (this BooksDatabaseImpl) GetBook(id string) (model.Book, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return model.Book{}, err
	}

	var book model.Book
	query := "SELECT * FROM " + tableName + " WHERE id = $1"

	err = this.DB.QueryRow(query, uuid).Scan(&book.Id, &book.Title, &book.Price)
	if err != nil {
		log.Println("No result to get specific book, err: ", err)
	}

	return book, err
}

func (this BooksDatabaseImpl) CreateBook(book model.Book) error {
	query := "INSERT INTO " + tableName + "(id, title, price) VALUES ($1, $2, $3)"
	_, err := this.DB.Exec(query, book.Id, book.Title, book.Price)
	if err != nil {
		return err
	}

	return nil
}

func (this BooksDatabaseImpl) DeleteBook(id string) error {
	query := "DELETE FROM " + tableName + " WHERE id = $1"
	res, err := this.DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Not found.")
	}

	return nil
}

func (this BooksDatabaseImpl) PutBook(book model.Book) error {
	book, err := this.GetBook(book.Id.String())
	if err != nil {
		return err
	}

	query := "UPDATE " + tableName + " SET title = $1, price = $2 WHERE id = $3"
	_, err = this.DB.Exec(query, book.Title, book.Price, book.Id)
	if err != nil {
		return err
	}

	return nil
}
