package model

import (
	"github.com/google/uuid"
)

type CreateComment struct {
	BookId  uuid.UUID `json:"book_id" validate:"required"`
	Comment string    `json:"comment" validate:"required,min=5,max=500"`
	Rating  int       `json:"rating" validate:"required,gte=1,lte=5"`
}

type UpdateComment struct {
	Comment string `json:"comment" validate:"required,min=5,max=500"`
	Rating  int    `json:"rating" validate:"required,gte=1,lte=5"`
}
