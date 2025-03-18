package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

type IUserRepository interface {
	GetUsers(users *[]entity.User, page, pageSize int) error
	GetUser(user *entity.User, userId uuid.UUID) error
	GetUserByEmail(email string) (user *entity.User, err error)
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
	if errors.Is(err, sql.ErrNoRows) {
		return &response.UserNotFound
	}
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (user *entity.User, err error) {
	query := `SELECT * FROM users WHERE email = $1`

	user = &entity.User{}
	err = r.db.Get(user, query, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &response.UserNotFound
	}

	return user, err
}
