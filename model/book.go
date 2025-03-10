package model

type BookReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type BookSearch struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	SearchParam string `json:"search_param"`
}
