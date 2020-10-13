package handler

import "net/http"

func HandleNoContent(w http.ResponseWriter) {
	http.Error(w, "No Content", http.StatusNoContent)
}

func HandleCreated(w http.ResponseWriter) {
	http.Error(w, "Created", http.StatusCreated)
}
