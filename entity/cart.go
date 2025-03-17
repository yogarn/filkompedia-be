package entity

import "github.com/google/uuid"

type Cart struct {
	CartId uuid.UUID `json:"cart_id" db:"cart_id"`
	UserId uuid.UUID `json:"user_id" db:"user_id"`
	BookId uuid.UUID `json:"book_id" db:"book_id"`
	Amount int       `json:"amount" db:"amount"`
}
