package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Author      string    `json:"author" db:"author"`
	ReleaseDate string    `json:"release_date" db:"release_date"`
	Price       float64   `json:"price" db:"price"`
}
