package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type IBookRepository interface {
	GetBooks(books *[]entity.Book, page, pageSize int) error
	SearchBooks(books *[]entity.Book, page, pageSize int, searchQuery string) error
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) IBookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetBooks(books *[]entity.Book, page, pageSize int) error {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := `SELECT * FROM books ORDER BY release_date DESC LIMIT $1 OFFSET $2`
	err := r.db.Select(&books, query, pageSize, offset)
	return err
}

func (r *BookRepository) SearchBooks(books *[]entity.Book, page, pageSize int, searchQuery string) error {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := `
		SELECT * FROM books 
		WHERE 
			title ILIKE $1 OR 
			author ILIKE $1 OR 
			description ILIKE $1 
		ORDER BY release_date DESC 
		LIMIT $2 OFFSET $3
	`

	searchPattern := "%" + searchQuery + "%"

	err := r.db.Select(&books, query, searchPattern, pageSize, offset)
	return err
}
