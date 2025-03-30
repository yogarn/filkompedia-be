package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

type ICheckoutRepository interface {
	GetUserCheckouts(userId uuid.UUID) (*[]entity.Checkout, error)
	GetCheckoutCarts(checkoutId uuid.UUID) (*[]entity.Cart, error)
	AddCheckoutId(cartID uuid.UUID, checkoutId uuid.UUID) error
	NewCheckout(userId, CheckoutId uuid.UUID) error
}

type CheckoutRepository struct {
	db *sqlx.DB
}

func NewCheckoutRepository(db *sqlx.DB) ICheckoutRepository {
	return &CheckoutRepository{
		db: db,
	}
}

func (r *CheckoutRepository) GetUserCheckouts(userId uuid.UUID) (*[]entity.Checkout, error) {
	var checkouts []entity.Checkout
	query := `SELECT * FROM checkouts WHERE user_id = $1`
	err := r.db.Select(&checkouts, query, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, &response.CheckoutNotFound
	}

	return &checkouts, err
}

func (r *CheckoutRepository) GetCheckoutCarts(checkoutId uuid.UUID) (*[]entity.Cart, error) {
	var carts []entity.Cart
	query := `SELECT * FROM carts WHERE checkout_id = $1`
	err := r.db.Select(&carts, query, checkoutId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, &response.CheckoutNotFound
	}

	return &carts, err
}

func (r *CheckoutRepository) AddCheckoutId(cartID uuid.UUID, checkoutId uuid.UUID) error {
	query := `UPDATE carts SET checkout_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, checkoutId, cartID)
	return err
}

func (r *CheckoutRepository) NewCheckout(CheckoutId, userId uuid.UUID) error {
	query := `INSERT INTO checkouts (id, user_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, CheckoutId, userId)
	return err
}
