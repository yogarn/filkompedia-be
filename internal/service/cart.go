package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/internal/repository"
)

type ICartService interface {
}

type CartService struct {
}

func NewCartService(db *sqlx.DB, cartRepo repository.ICartRepository) {

}
