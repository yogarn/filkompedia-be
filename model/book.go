package model

import "github.com/google/uuid"

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
	Title        string  `json:"title" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Introduction string  `json:"introduction" validate:"required"`
	Image        string  `json:"image"`
	File         string  `jsob:"file"`
	Author       string  `json:"author" validate:"required"`
	ReleaseDate  string  `json:"release_date" validate:"required"`
	Price        float64 `json:"price"`
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
