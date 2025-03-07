package service

import (
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
)

type Service struct {
	UserService IUserService
	AuthService IAuthService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt) *Service {
	return &Service{
		UserService: NewUserService(repository),
		AuthService: NewAuthService(repository.AuthRepository, bcrypt),
	}
}
