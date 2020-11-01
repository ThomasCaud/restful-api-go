package db

import (
	"database/sql"
	"fmt"
	"os"
)

// DB contain direct access to the DB
type DB struct {
	*sql.DB
}

// InitializeDatabaseConnection try to connect to database
// Using env vars
func InitializeDatabaseConnection() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open(os.Getenv("DB_TYPE"), dbinfo)

	return db, err
}
