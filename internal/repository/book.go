package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type IBookRepository interface {
	GetBooks(books *[]entity.Book, page, pageSize int) error
	SearchBooks(books *[]entity.Book, page, pageSize int, searchQuery string) error
	GetBook(book *entity.Book, bookId uuid.UUID) error
	CreateBook(book *entity.Book) error
	DeleteBook(bookId uuid.UUID) error
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
	err := r.db.Select(books, query, pageSize, offset)
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
		LIMIT $2 OFFSET $3`
	searchPattern := "%" + searchQuery + "%"
	err := r.db.Select(books, query, searchPattern, pageSize, offset)
	return err
}

func (r *BookRepository) GetBook(book *entity.Book, bookId uuid.UUID) error {
	query := `SELECT * FROM books WHERE id = $1 LIMIT 1`
	err := r.db.Get(book, query, bookId)
	return err
}

func (r *BookRepository) CreateBook(book *entity.Book) error {
	query := `INSERT INTO books (id, title, description, author, release_date, price) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, book.Id, book.Title, book.Description, book.Author, book.ReleaseDate, book.Price)
	return err
}

func (r *BookRepository) DeleteBook(bookId uuid.UUID) error {
	query := `DELETE FROM books WHERE id = $1`
	result, err := r.db.Exec(query, bookId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
