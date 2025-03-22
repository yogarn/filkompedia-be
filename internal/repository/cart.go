package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
)

type ICartRepository interface {
	GetUserCart(carts *[]entity.Cart, user *entity.User) error
	GetCart(cart *entity.Cart, cartId uuid.UUID) error
	AddToCart(user *entity.User, book *entity.Book, amount int) error
	RemoveFromCart(cartId uuid.UUID) error
}

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) ICartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) GetUserCart(carts *[]entity.Cart, user *entity.User) error {
	query := `SELECT * FROM carts WHERE user_id = $1`
	err := r.db.Select(carts, query, user.Id)
	return err
}

func (r *CartRepository) GetCart(cart *entity.Cart, cartId uuid.UUID) error {
	query := `SELECT * FROM carts WHERE id = $1`
	err := r.db.Get(cart, query, cartId)
	return err
}

func (r *CartRepository) AddToCart(user *entity.User, book *entity.Book, amount int) error {
	if amount < 1 {
		return errors.New("invalid amount")
	}

	var cart entity.Cart
	if r.doesCartExist(&cart, user.Id, book.Id); cart.Amount > 0 {
		query := `UPDATE carts SET amount = $1 + $2 WHERE user_id = $3 AND book_id = $4 `
		_, err := r.db.Exec(query, cart.Amount, amount, user.Id, book.Id)
		return err
	}

	query := `INSERT INTO carts (id, user_id, book_id, amount) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, uuid.New(), user.Id, book.Id, amount)
	return err
}

func (r *CartRepository) RemoveFromCart(cartId uuid.UUID) error {
	query := `DELETE FROM carts WHERE id = $1`
	_, err := r.db.Exec(query, cartId)
	return err
}

func (r *CartRepository) doesCartExist(cart *entity.Cart, userId uuid.UUID, bookId uuid.UUID) error {
	query := `SELECT * FROM carts WHERE user_id = $1 AND book_id = $2 LIMIT 1`
	err := r.db.Get(cart, query, userId, bookId)
	return err
}
