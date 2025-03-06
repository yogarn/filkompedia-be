package repository

import "github.com/jmoiron/sqlx"

type IUserRepository interface {
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
