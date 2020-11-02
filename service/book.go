package service

import (
	"errors"

	"github.com/ThomasCaud/go-rest-api/model"
	store "github.com/ThomasCaud/go-rest-api/store/postgres"
)

// GetCollection returns books collection
func GetCollection(store store.BooksDatabaseImpl) ([]*model.Book, error) {
	books, err := store.GetBooks()
	if err != nil {
		return nil, errors.New("error while getting books")
	}

	return books, nil
}
