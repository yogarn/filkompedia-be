package service

import (
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
)

type Service struct {
	UserService IUserService
	AuthService IAuthService
	BookService IBookService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
		AuthService: NewAuthService(repository.AuthRepository, bcrypt),
		BookService: NewBookService(repository.BookRepository),
	}
}
