package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type IBookService interface {
	GetBooks(books *[]entity.Book, bookReq model.BookReq) error
	SearchBooks(books *[]entity.Book, bookSearch model.BookSearch) error
	CreateBook(create *model.CreateBook) error
}

type BookService struct {
	bookRepo repository.IBookRepository
}

func NewBookService(bookRepo repository.IBookRepository) IBookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (s *BookService) GetBooks(books *[]entity.Book, bookReq model.BookReq) error {
	return s.bookRepo.GetBooks(books, bookReq.Page, bookReq.PageSize)
}

func (s *BookService) SearchBooks(books *[]entity.Book, bookSearch model.BookSearch) error {
	return s.bookRepo.SearchBooks(books, bookSearch.Page, bookSearch.PageSize, bookSearch.SearchParam)
}
func (s *BookService) CreateBook(create *model.CreateBook) error {
	return s.bookRepo.CreateBook(&entity.Book{
		Id:          uuid.New(),
		Title:       create.Title,
		Description: create.Description,
		Author:      create.Author,
		ReleaseDate: create.ReleaseDate,
		Price:       create.Price,
	})
}
