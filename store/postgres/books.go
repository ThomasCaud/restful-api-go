package postgres

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ThomasCaud/go-rest-api/model"
)

// BooksDatabaseImpl contain direct access to the DB
type BooksDatabaseImpl struct {
	DB *sql.DB
}

// GetBooks return all books
func (dbImpl BooksDatabaseImpl) GetBooks() ([]*model.Book, error) {
	rows, err := dbImpl.DB.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	defer rows.Close()

	books := make([]*model.Book, 0)
	for rows.Next() {
		book := new(model.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Price)

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

// GetBook return book with the parameter id, error otherwise
func (dbImpl BooksDatabaseImpl) GetBook(id string) (*model.Book, error) {
	query := "SELECT * FROM books WHERE id = $1 LIMIT 100"

	book := model.Book{}
	err := dbImpl.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Price)
	if err != nil {
		log.Println("No result to get specific book, err: ", err)
	}

	return &book, err
}

// CreateBook create book using book parameter
func (dbImpl BooksDatabaseImpl) CreateBook(book model.Book) error {
	query := "INSERT INTO books (id, title, price) VALUES ($1, $2, $3)"
	_, err := dbImpl.DB.Exec(query, book.ID, book.Title, book.Price)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBook delete the book corresponding to the giving ID parameter
func (dbImpl BooksDatabaseImpl) DeleteBook(id string) error {
	query := "DELETE FROM books WHERE id = $1"
	res, err := dbImpl.DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Not found")
	}

	return nil
}

// PutBook update the book
func (dbImpl BooksDatabaseImpl) PutBook(book model.Book) error {
	_, err := dbImpl.GetBook(book.ID.String())

	if err != nil {
		return err
	}

	query := "UPDATE books SET title = $1, price = $2 WHERE id = $3"
	_, err = dbImpl.DB.Exec(query, book.Title, book.Price, book.ID)
	if err != nil {
		return err
	}

	return nil
}
