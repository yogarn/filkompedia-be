package service

import "github.com/yogarn/filkompedia-be/internal/repository"

type IUserService interface {
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}
