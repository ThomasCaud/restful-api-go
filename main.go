package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ThomasCaud/go-rest-api/handler"
	"github.com/ThomasCaud/go-rest-api/store"
	"github.com/ThomasCaud/go-rest-api/store/postgres"
)

func main() {
	database, err := store.InitializeDatabaseConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	} else {
		log.Println("Database connection succeed")
	}

	// Wait for db container to be ready
	// todo wait via docker-compose
	// Mon bin ne devrait pas être responsable de la gestion des migrations
	// A supprimer et réutiliser ancienne version
	// Les inserts: les faire via l'appel API (simple CURL en sh, python, go...)
	time.Sleep(3 * time.Second)
	err = store.ExecuteMigrations(database)
	if err != nil {
		log.Fatal("Migrations failed: ", err.Error())
	}

	var booksDbImpl = postgres.BooksDatabaseImpl{}
	booksDbImpl.DB = database

	app := &handler.App{
		BooksDatabase: booksDbImpl,
	}

	log.Fatal(http.ListenAndServe(":3000", handler.GetRouter(app)))
}
