package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type IUserRepository interface {
	GetUsers(users *[]entity.User, page, pageSize int) error
	GetUser(user *entity.User, userId uuid.UUID) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers(users *[]entity.User, page, pageSize int) error {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := `SELECT * FROM users ORDER BY id ASC LIMIT $1 OFFSET $2`
	err := r.db.Select(users, query, pageSize, offset)
	return err
}

func (r *UserRepository) GetUser(user *entity.User, userId uuid.UUID) error {
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.Get(user, query, userId)
	return err
}
