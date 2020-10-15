package model

type Book struct {
	Id    string `json:"Id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
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
