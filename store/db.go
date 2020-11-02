package store

import (
	"database/sql"
	"fmt"
	"os"

	// postgres driver
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

// DB contain direct access to the DB
type DB struct {
	*sql.DB
}

// ExecuteMigrations get and execute migrations
func ExecuteMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "store/migrations",
	}

	n, err := migrate.Exec(db, os.Getenv("DB_TYPE"), migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}

// InitializeDatabaseConnection try to connect to database
// Using env vars
func InitializeDatabaseConnection() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open(os.Getenv("DB_TYPE"), dbinfo)

	return db, err
}
