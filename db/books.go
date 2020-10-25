package db

import (
	"database/sql"
	"log"

	"github.com/ThomasCaud/go-rest-api/model"
)

type BooksDatabaseImpl struct {
	DB *sql.DB
}

func (this BooksDatabaseImpl) GetBooks() ([]*model.Book, error) {
	rows, err := this.DB.Query("select * from books")
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

func (this BooksDatabaseImpl) GetBook(id int) (model.Book, error) {
	var book model.Book
	query := "select * from books where id = $1"

	err := this.DB.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Price)
	if err != nil {
		log.Println("No result to get specific book, err: %s", err)
	}

	return book, err
}
