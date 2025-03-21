package service

import (
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
	"github.com/yogarn/filkompedia-be/pkg/smtp"
)

type Service struct {
	UserService IUserService
	AuthService IAuthService
	BookService IBookService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt, smtp *smtp.SMTPClient) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
		AuthService: NewAuthService(repository.AuthRepository, repository.UserRepository, bcrypt, jwt, smtp),
		BookService: NewBookService(repository.BookRepository),
	}
}
