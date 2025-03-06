package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UserRepository IUserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
