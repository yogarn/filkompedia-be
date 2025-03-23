package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type IBookService interface {
	SearchBooks(bookSearch model.BookSearch) (*[]model.BookResponse, error)
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

func (s *BookService) SearchBooks(bookSearch model.BookSearch) (*[]model.BookResponse, error) {
	var booksEntity []entity.Book
	err := s.bookRepo.SearchBooks(&booksEntity, bookSearch.Page, bookSearch.PageSize, bookSearch.SearchParam)
	if err != nil {
		return nil, err
	}

	booksResponse := make([]model.BookResponse, len(booksEntity))
	for i, book := range booksEntity {
		booksResponse[i] = model.BookResponse{
			Id:           book.Id,
			Title:        book.Title,
			Image:        book.Image,
			Description:  book.Description,
			Introduction: book.Introduction,
			Author:       book.Author,
			ReleaseDate:  book.ReleaseDate,
			Price:        book.Price,
		}
	}

	return &booksResponse, nil
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
