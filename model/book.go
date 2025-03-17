package model

import (
	"time"
)

type BookReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type BookSearch struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	SearchParam string `json:"search_param"`
}

type CreateBook struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"release_date"`
	Price       float64   `json:"price"`
}
