package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type ICheckoutService interface {
	GetUserCheckouts(userId uuid.UUID) ([]entity.Checkout, error)
	GetCheckoutCarts(checkoutId uuid.UUID) ([]entity.Cart, error)
	Checkout(checkoutReq model.CheckoutRequest) error
}

type CheckoutService struct {
	checkoutRepo repository.ICheckoutRepository
	cartRepo     repository.ICartRepository
}

func NewCheckoutService(checkoutRepo repository.ICheckoutRepository, cartRepo repository.ICartRepository) ICheckoutService {
	return &CheckoutService{
		checkoutRepo: checkoutRepo,
	}
}

func (s *CheckoutService) GetUserCheckouts(userId uuid.UUID) ([]entity.Checkout, error) {
	return s.checkoutRepo.GetUserCheckouts(userId)
}

func (s *CheckoutService) GetCheckoutCarts(checkoutId uuid.UUID) ([]entity.Cart, error) {
	return s.checkoutRepo.GetCheckoutCarts(checkoutId)
}

func (s *CheckoutService) Checkout(checkoutReq model.CheckoutRequest) error {
	for _, cart_id := range checkoutReq.CartsId {
		var cart entity.Cart
		if err := s.cartRepo.GetCart(&cart, cart_id); err != nil {
			return err
		}
		if cart.UserId != checkoutReq.UserId {
			return errors.New("invalid input at " + cart_id.String())
		}
	}

	newCheckoutId := uuid.New()
	for _, cart_id := range checkoutReq.CartsId {
		if err := s.checkoutRepo.AddCheckoutId(cart_id, newCheckoutId); err != nil {
			return err
		}
	}

	return nil
}
