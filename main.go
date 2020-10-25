package main

import (
	"log"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/db"
	"github.com/ThomasCaud/go-rest-api/handler"
)

func main() {
	database, err := db.InitializeDatabaseConnection()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	} else {
		log.Println("Database connection succeed")
	}

	var booksDbImpl = db.BooksDatabaseImpl{}
	booksDbImpl.DB = database

	app := &handler.App{
		BooksDatabase: booksDbImpl,
	}

	log.Fatal(http.ListenAndServe(":3000", handler.GetRouter(app)))
}
