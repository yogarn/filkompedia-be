package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type IAuthRepository interface {
	Register(user *entity.User) (err error)
	Login(session *entity.Session) (err error)
	GetSessions(userId uuid.UUID) (sessions *[]entity.Session, err error)
}

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(user *entity.User) (err error) {
	query := `INSERT INTO users (id, username, email, password, role_id) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(query, user.Id, user.Username, user.Email, user.Password, user.RoleId)
	return err
}

func (r *AuthRepository) Login(session *entity.Session) (err error) {
	query := `
		INSERT INTO sessions (user_id, token, ip_address, expires_at, user_agent, device_id)
		VALUES (:user_id, :token, :ip_address, :expires_at, :user_agent, :device_id)
	`

	_, err = r.db.NamedExec(query, session)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetSessions(userId uuid.UUID) (sessions *[]entity.Session, err error) {
	query := `
		SELECT * FROM sessions WHERE user_id = $1
	`

	sessions = &[]entity.Session{}

	err = r.db.Select(sessions, query, userId)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
