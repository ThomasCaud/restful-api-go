package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/ThomasCaud/go-rest-api/handler"
)

func main() {
	database, err := db.InitializeDatabaseConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	} else {
		log.Println("Database connection succeed")
	}

	// Wait for db container to be ready
	// todo wait via docker-compose
	time.Sleep(3 * time.Second)
	err = db.ExecuteMigrations(database)
	if err != nil {
		log.Fatal("Migrations failed: ", err.Error())
	}

	var booksDbImpl = db.BooksDatabaseImpl{}
	booksDbImpl.DB = database

	app := &handler.App{
		BooksDatabase: booksDbImpl,
	}

	log.Fatal(http.ListenAndServe(":3000", handler.GetRouter(app)))
}
