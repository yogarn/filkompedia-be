package service

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

type IAuthService interface {
	Register(registerReq *model.RegisterReq) (user *entity.User, err error)
	Login(loginReq *model.LoginReq, ipAddress string, userAgent string, expiry int) (loginRes *model.LoginRes, err error)
}

type AuthService struct {
	AuthRepository repository.IAuthRepository
	UserRepository repository.IUserRepository
	Bcrypt         bcrypt.IBcrypt
	Jwt            jwt.IJwt
}

func NewAuthService(authRepository repository.IAuthRepository, userRepository repository.IUserRepository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
		UserRepository: userRepository,
		Bcrypt:         bcrypt,
		Jwt:            jwt,
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

func (s *AuthService) Login(loginReq *model.LoginReq, ipAddress string, userAgent string, expiry int) (loginRes *model.LoginRes, err error) {
	user, err := s.UserRepository.GetUserByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	err = s.Bcrypt.CompareAndHashPassword(user.Password, loginReq.Password)
	if err != nil {
		return nil, &response.InvalidCredentials
	}

	token, err := s.Jwt.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRandomString(32)
	if err != nil {
		return nil, err
	}

	session := &entity.Session{
		UserId:    user.Id,
		Token:     refreshToken,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		DeviceId:  generateDeviceID(),
		ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Second),
	}

	err = s.AuthRepository.Login(session)
	if err != nil {
		return nil, err
	}

	return &model.LoginRes{
		JwtToken:     token,
		RefreshToken: refreshToken,
	}, nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func generateDeviceID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
