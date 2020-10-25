package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

const (
	DB_USER     = "user"
	DB_PASSWORD = "pass"
	DB_NAME     = "bookstore"
)

func InitializeDatabaseConnection() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, err
	}

	return db, nil
}
