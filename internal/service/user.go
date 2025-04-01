package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type IUserService interface {
	GetProfiles(profiles *[]model.Profile, profilesReq model.ProfilesReq) error
	GetProfile(profile *model.Profile, userId uuid.UUID) error
	GetUserById(user *entity.User, userId uuid.UUID) (err error)
	UpdateRole(userProfile *model.RoleUpdate) error
	EditProfile(edit *model.EditProfile) error
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetProfiles(profiles *[]model.Profile, profilesReq model.ProfilesReq) error {
	var users []entity.User
	if err := s.UserRepository.GetUsers(&users, profilesReq.Page, profilesReq.PageSize); err != nil {
		return err
	}

	*profiles = make([]model.Profile, len(users))
	for i, user := range users {
		(*profiles)[i] = model.UserToProfile(user)
	}

	return nil
}

func (s *UserService) GetUserById(user *entity.User, userId uuid.UUID) (err error) {
	err = s.UserRepository.GetUser(user, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetProfile(profile *model.Profile, userId uuid.UUID) error {
	var user entity.User
	if err := s.UserRepository.GetUser(&user, userId); err != nil {
		return err
	}

	*profile = model.UserToProfile(user)

	return nil
}

func (s *UserService) UpdateRole(userProfile *model.RoleUpdate) error {
	return s.UserRepository.UpdateRole(userProfile.Id, userProfile.RoleId)
}

func (s *UserService) EditProfile(edit *model.EditProfile) error {
	var user entity.User
	if err := s.UserRepository.GetUser(&user, edit.Id); err != nil {
		return err
	}

	//todo improve this
	if edit.Username == "" {
		edit.Username = user.Username
	}
	if edit.Email == "" {
		edit.Email = user.Email
	}

	if err := s.UserRepository.EditUser(edit); err != nil {
		return err
	}

	return nil
}
