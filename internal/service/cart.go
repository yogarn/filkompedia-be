package service

// import (
// 	"github.com/google/uuid"
// 	"github.com/yogarn/filkompedia-be/entity"
// 	"github.com/yogarn/filkompedia-be/internal/repository"
// 	"github.com/yogarn/filkompedia-be/model"
// )

// type ICartService interface {
// 	GetUserCart(carts *[]entity.Cart, UserId uuid.UUID) error
// 	GetCart(param *model.CartParam, cart *entity.Cart) error
// 	AddToCart(add model.AddToCart) error
// 	RemoveFromCart(param *model.CartParam) error
// }

// type CartService struct {
// 	cartRepo repository.ICartRepository
// }

// func NewCartService(cartRepo repository.ICartRepository) ICartService {
// 	return &CartService{
// 		cartRepo: cartRepo,
// 	}
// }

// func (s *CartService) GetUserCart(carts *[]entity.Cart, UserId uuid.UUID) error {
// 	return s.cartRepo.GetUserCart(carts, UserId)
// }

// func (s *CartService) GetCart(param *model.CartParam, cart *entity.Cart) error {

// }

// func (s *CartService) AddToCart(add model.AddToCart) error {

// }

// func (s *CartService) RemoveFromCart(param *model.CartParam) error {

// }
