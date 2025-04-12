package model

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
)

type ProfilesReq struct {
	Page     int `json:"page" validate:"required,min=1"`
	PageSize int `json:"page_size" validate:"required,min=1"`
}

type Profile struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	RoleId   int       `json:"roleId" db:"role_id"`
}

type RoleUpdate struct {
	Id     uuid.UUID `json:"id" db:"id" validate:"required"`
	RoleId int       `json:"roleId" db:"role_id" validate:"required,min=0,max=1"`
}

type EditProfile struct {
	Id         uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
	Username   string    `json:"username" db:"username" validate:"required,lte=32"`
	IsVerified bool      `db:"is_verified"`
}

type ProfilePicture struct {
	File *multipart.FileHeader `form:"file" validate:"required,image_type,image_size"`
}

func UserToProfile(user entity.User) Profile {
	return Profile{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RoleId:   user.RoleId,
	}
}
