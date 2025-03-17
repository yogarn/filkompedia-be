package model

type BookReq struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page_size" validate:"required"`
}

type BookSearch struct {
	Page        int    `json:"page" validate:"required"`
	PageSize    int    `json:"page_size" validate:"required"`
	SearchParam string `json:"search_param" validate:"required"`
}

type CreateBook struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	ReleaseDate string  `json:"release_date"`
	Price       float64 `json:"price"`
}
