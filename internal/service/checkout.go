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
	Checkout(checkoutReq model.CheckoutRequest, checkoutId uuid.UUID) (float64, error)
}

type CheckoutService struct {
	checkoutRepo repository.ICheckoutRepository
	cartRepo     repository.ICartRepository
	bookRepo     repository.IBookRepository
}

func NewCheckoutService(checkoutRepo repository.ICheckoutRepository, cartRepo repository.ICartRepository, bookRepo repository.IBookRepository) ICheckoutService {
	return &CheckoutService{
		checkoutRepo: checkoutRepo,
		cartRepo:     cartRepo,
		bookRepo:     bookRepo,
	}
}

func (s *CheckoutService) GetUserCheckouts(userId uuid.UUID) ([]entity.Checkout, error) {
	return s.checkoutRepo.GetUserCheckouts(userId)
}

func (s *CheckoutService) GetCheckoutCarts(checkoutId uuid.UUID) ([]entity.Cart, error) {
	return s.checkoutRepo.GetCheckoutCarts(checkoutId)
}

func (s *CheckoutService) Checkout(checkoutReq model.CheckoutRequest, checkoutId uuid.UUID) (totalPrice float64, err error) {
	for _, cart_id := range checkoutReq.CartsId {
		var cart entity.Cart
		if err := s.cartRepo.GetCart(&cart, cart_id); err != nil {
			return 0, err
		}

		var book entity.Book
		if err := s.bookRepo.GetBook(&book, cart.BookId); err != nil {
			return 0, err
		}

		if cart.UserId != checkoutReq.UserId {
			return 0, errors.New("invalid input at " + cart_id.String())
		}

		totalPrice += (float64(cart.Amount) * book.Price)
	}

	if err := s.checkoutRepo.NewCheckout(checkoutId, checkoutReq.UserId); err != nil {
		return 0, err
	}

	for _, cart_id := range checkoutReq.CartsId {
		if err := s.checkoutRepo.AddCheckoutId(cart_id, checkoutId); err != nil {
			return 0, err
		}
	}

	return totalPrice, nil
}
