package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
)

type IAuthService interface {
	Register(registerReq *model.RegisterReq) (user *entity.User, err error)
}

type AuthService struct {
	AuthRepository repository.IAuthRepository
	Bcrypt         bcrypt.IBcrypt
}

func NewAuthService(authRepository repository.IAuthRepository, bcrypt bcrypt.IBcrypt) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
		Bcrypt:         bcrypt,
	}
}

func (s *AuthService) Register(registerReq *model.RegisterReq) (user *entity.User, err error) {
	hashedpassword, err := s.Bcrypt.GenerateFromPassword(registerReq.Password)
	if err != nil {
		return nil, err
	}

	user = &entity.User{
		Id:       uuid.New(),
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: hashedpassword,
		RoleId:   0,
	}

	err = s.AuthRepository.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
