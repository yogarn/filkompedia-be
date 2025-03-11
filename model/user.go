package model

import (
	"github.com/yogarn/filkompedia-be/entity"
)

type ProfilesReq struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page_size" validate:"required"`
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
