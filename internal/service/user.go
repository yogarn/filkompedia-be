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
	DeleteUser(userId uuid.UUID) error
}

type UserService struct {
	UserRepository     repository.IUserRepository
	CartRepository     repository.ICartRepository
	PaymentRepository  repository.IPaymentRepository
	AuthRepository     repository.IAuthRepository
	CheckoutRepository repository.ICheckoutRepository
}

func NewUserService(userRepository repository.IUserRepository, cartRepository repository.ICartRepository, paymentRepository repository.IPaymentRepository, authRepository repository.IAuthRepository, checkoutRepository repository.ICheckoutRepository) IUserService {
	return &UserService{
		UserRepository:     userRepository,
		CartRepository:     cartRepository,
		PaymentRepository:  paymentRepository,
		AuthRepository:     authRepository,
		CheckoutRepository: checkoutRepository,
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
	err := s.UserRepository.UpdateRole(userProfile.Id, userProfile.RoleId)
	if err != nil {
		return err
	}

	return nil
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

	if err := s.UserRepository.EditUser(edit); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(userId uuid.UUID) error {
	var user entity.User
	if err := s.UserRepository.GetUser(&user, userId); err != nil {
		return err
	}

	//todo implement transaction

	if err := s.AuthRepository.ClearToken(userId); err != nil {
		return err
	}

	if err := s.PaymentRepository.DeleteUser(userId); err != nil {
		return err
	}

	if err := s.CartRepository.DeleteUserCart(userId); err != nil {
		return err
	}

	if err := s.CartRepository.DeleteUser(userId); err != nil {
		return err
	}

	if err := s.CheckoutRepository.DeleteUser(userId); err != nil {
		return err
	}

	if err := s.UserRepository.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
