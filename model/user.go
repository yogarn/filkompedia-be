package model

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
)

type ProfilesReq struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"page_size" validate:"required"`
}

type Profile struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	RoleId   int       `json:"roleId" db:"role_id"`
}

func UserToProfile(user entity.User) Profile {
	return Profile{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.RoleId,
	}
}
