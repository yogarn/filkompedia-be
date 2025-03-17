package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id          uuid.UUID `json:"book_id" db:"book_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Author      string    `json:"author" db:"author"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Price       float64   `json:"price" db:"price"`
}
