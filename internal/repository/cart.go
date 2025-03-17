package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type ICartRepository interface {
	GetUserCart(carts *[]entity.Cart, userId uuid.UUID) error
	GetCart(cart *entity.Cart, cartId uuid.UUID) error
	AddToCart(userId uuid.UUID, bookId uuid.UUID, amount int) error
	RemoveFromCart(cartId uuid.UUID) error
}

type CartRepository struct {
	db       *sqlx.DB
	userRepo IUserRepository
	bookRepo IBookRepository
}

func NewCartRepositorcleary(db *sqlx.DB, userRepo IUserRepository, bookRepo IBookRepository) ICartRepository {
	return &CartRepository{
		db:       db,
		userRepo: userRepo,
		bookRepo: bookRepo, 
	}
}

func (r *CartRepository) GetUserCart(carts *[]entity.Cart, userId uuid.UUID) error {
	var user entity.User
	if err := r.userRepo.GetUser(&user, userId); err != nil {
		return err
	}

	query := `SELECT * FROM carts WHERE user_id = $1`
	err := r.db.Select(carts, query, userId)
	return err
}

func (r *CartRepository) GetCart(cart *entity.Cart, cartId uuid.UUID) error {
	query := `SELECT * FROM carts WHERE cart_id = $1`
	err := r.db.Select(cart, query, cartId)
	return err
}

func (r *CartRepository) AddToCart(userId uuid.UUID, bookId uuid.UUID, amount int) error {
	var user entity.User
	if err := r.userRepo.GetUser(&user, userId); err != nil {
		return err
	}

	var book entity.Book
	if err := r.bookRepo.GetBook(&book, bookId); err != nil {
		return err
	}

	if amount < 1 {
		return errors.New("invalid amount")
	}

	var cart entity.Cart
	if r.doesCartExist(&cart, userId, bookId); cart.Amount > 0 {
		query := `UPDATE carts SET amount = $1 + $2 WHERE user_id = $3 AND book_id = $4 `
		_, err := r.db.Exec(query, cart.Amount, amount, userId, bookId)
		return err
	}

	query := `INSERT INTO carts (cart_id, user_id, book_id, amount) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, uuid.New(), userId, bookId, amount)
	return err
}

func (r *CartRepository) RemoveFromCart(cartId uuid.UUID) error {
	query := `DELETE FROM carts WHERE cart_id = $1`
	_, err := r.db.Exec(query, cartId)
	return err
}

func (r *CartRepository) doesCartExist(cart *entity.Cart, userId uuid.UUID, bookId uuid.UUID) error {
	query := `SELECT TOP 1 FROM carts WHERE user_id = $1 AND book_id = $2`
	err := r.db.Select(cart, query, userId, bookId)
	return err
}
