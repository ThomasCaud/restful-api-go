package main

import (
	"log"
	"net/http"

	"github.com/ThomasCaud/go-rest-api/handler"
)

func handleRequests() {
	log.Fatal(http.ListenAndServe(":3000", getRouter()))
}

func main() {
	handler.PopulateBooks()
	handleRequests()
}
