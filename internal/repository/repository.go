package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository IUserRepository
	AuthRepository IAuthRepository
	BookRepository IBookRepository
	CartRepository ICartRepository
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		AuthRepository: NewAuthRepository(db, redis),
		BookRepository: NewBookRepository(db),
		CartRepository: NewCartRepository(db),
	}
}
