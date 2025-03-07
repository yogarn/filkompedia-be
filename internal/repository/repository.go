package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UserRepository IUserRepository
	AuthRepository IAuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		AuthRepository: NewAuthRepository(db),
	}
}
