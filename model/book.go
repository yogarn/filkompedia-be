package model

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
)

type BookReq struct {
	Page     int `json:"page" validate:"required,min=1"`
	PageSize int `json:"page_size" validate:"required,min=1"`
}

type BookSearch struct {
	Page        int    `json:"page" validate:"required,min=1"`
	PageSize    int    `json:"page_size" validate:"required,min=1"`
	SearchParam string `json:"search_param" validate:"required,min=1"`
}

type CreateBook struct {
	Title        string  `json:"title" validate:"required,gte=5"`
	Description  string  `json:"description" validate:"required,gte=10"`
	Introduction string  `json:"introduction" validate:"required,gte=10"`
	Image        string  `json:"image" validate:"required,url"`
	File         string  `jsob:"file"`
	Author       string  `json:"author" validate:"required,gte=5"`
	ReleaseDate  string  `json:"release_date" validate:"required,rfc3339date"`
	Price        float64 `json:"price" validate:"required,min=1000"`
}

type BookResponse struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Introduction string    `json:"introduction"`
	Image        string    `json:"image"`
	Author       string    `json:"author"`
	ReleaseDate  string    `json:"release_date"`
	Price        float64   `json:"price"`
}

type EditBook struct {
	Id           uuid.UUID `json:"id" db:"id" validate:"required"`
	Title        string    `json:"title" db:"title" validate:"omitempty,gte=5"`
	Description  string    `json:"description" db:"description" validate:"omitempty,gte=10"`
	Introduction string    `json:"introduction" db:"introduction" validate:"omitempty,gte=10"`
	Image        string    `json:"image" db:"image" validate:"omitempty,url"`
	Author       string    `json:"author" db:"author" validate:"omitempty,gte=5"`
	ReleaseDate  string    `json:"release_date" db:"release_date" validate:"omitempty,rfc3339date"` //todo make a validator for date
	Price        float64   `json:"price" db:"price" validate:"omitempty,min=1000"`
}

type BookCover struct {
	File *multipart.FileHeader `form:"file" validate:"required,image_type,image_size"`
}

func BookToBookResponse(book entity.Book) BookResponse {
	return BookResponse{
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
