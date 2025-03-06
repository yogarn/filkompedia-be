package service

import "github.com/yogarn/filkompedia-be/internal/repository"

type Service struct {
	UserService IUserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repository),
	}
}
