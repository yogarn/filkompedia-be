package model

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
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	ReleaseDate string  `json:"release_date"`
	Price       float64 `json:"price"`
}
