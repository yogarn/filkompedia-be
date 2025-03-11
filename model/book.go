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
