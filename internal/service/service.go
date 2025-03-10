package service

import (
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
)

type Service struct {
	UserService IUserService
	AuthService IAuthService
	BookService IBookService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
		AuthService: NewAuthService(repository.AuthRepository, repository.UserRepository, bcrypt, jwt),
		BookService: NewBookService(repository.BookRepository),
	}
}
