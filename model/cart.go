package model

import "github.com/google/uuid"

type CartParam struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}

type AddToCart struct {
	BookId uuid.UUID `json:"book_id"`
	UserId uuid.UUID `json:"user_id"`
	Amount int       `json:"amount"`
}
