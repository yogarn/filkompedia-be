package model

import "github.com/google/uuid"

type CheckoutRequest struct {
	UserId  uuid.UUID   `json:"user_id" validate:"required,uuid"`
	CartsId []uuid.UUID `json:"carts_id" validate:"required"`
}
