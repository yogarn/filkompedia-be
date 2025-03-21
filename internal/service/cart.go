package service

import (
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type ICartService interface {
	GetUserCart(carts *[]entity.Cart, UserId uuid.UUID) error
	GetCart(cart *entity.Cart, cartId uuid.UUID) error
	AddToCart(add model.AddToCart) error
	RemoveFromCart(cartId uuid.UUID) error
}

type CartService struct {
	cartRepo repository.ICartRepository
	userRepo repository.IUserRepository
	bookRepo repository.IBookRepository
}

func NewCartService(cartRepo repository.ICartRepository, userRepo repository.IUserRepository, bookRepo repository.IBookRepository) ICartService {
	return &CartService{
		cartRepo: cartRepo,
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}

func (s *CartService) GetUserCart(carts *[]entity.Cart, UserId uuid.UUID) error {
	var user entity.User
	if err := s.userRepo.GetUser(&user, UserId); err != nil {
		return err
	}

	return s.cartRepo.GetUserCart(carts, &user)
}

func (s *CartService) GetCart(cart *entity.Cart, cartId uuid.UUID) error {
	return s.cartRepo.GetCart(cart, cartId)
}

func (s *CartService) AddToCart(add model.AddToCart) error {
	var user entity.User
	if err := s.userRepo.GetUser(&user, add.UserId); err != nil {
		return err
	}

	var book entity.Book
	if err := s.bookRepo.GetBook(&book, add.BookId); err != nil {
		return err
	}

	return s.cartRepo.AddToCart(&user, &book, add.Amount)
}

func (s *CartService) RemoveFromCart(cartId uuid.UUID) error {
	return s.cartRepo.RemoveFromCart(cartId)
}
