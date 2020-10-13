package main

import "net/http"

func handleNoContent(w http.ResponseWriter) {
	http.Error(w, "No Content", http.StatusNoContent)
}

func handleCreated(w http.ResponseWriter) {
	http.Error(w, "Created", http.StatusCreated)
}
