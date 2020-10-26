package model

import "github.com/google/uuid"

type Book struct {
	Id    uuid.UUID `json:"uuid"`
	Title string    `json:"title"`
	Price int       `json:"price"`
}
