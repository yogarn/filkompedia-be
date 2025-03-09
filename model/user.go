package model

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
)

type RegisterReq struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type ProfileReq struct {
	Id uuid.UUID `json:"id" db:"id"`
}

type ProfilesReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type Profile struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

func UserToProfile(user entity.User) Profile {
	return Profile{
		Username: user.Username,
		Email:    user.Email,
	}
}
