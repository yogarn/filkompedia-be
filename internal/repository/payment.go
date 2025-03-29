package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

type IPaymentRepository interface {
	GetPayment(paymentId uuid.UUID) (*entity.Payment, error)
	CreatePayment(payment entity.Payment) error
	UpdatePaymentStatus(statusId int, paymentId uuid.UUID) error
	//CheckUserBookPurchase()
}

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) IPaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) GetPayment(paymentId uuid.UUID) (*entity.Payment, error) {
	var payment entity.Payment
	query := `SELECT * FROM payments WHERE id = $1 LIMIT 1`
	err := r.db.Get(payment, query, paymentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.PaymentNotFound
		}
		return nil, err
	}
	return &payment, err
}

func (r *PaymentRepository) CreatePayment(payment entity.Payment) error {
	query := `
		INSERT INTO payments (id, token, user_id, checkout_id, total_price, status_id, created_at)
		VALUES (:id, :token, :user_id, :checkout_id, :total_price, :status_id, :created_at) 
	`
	_, err := r.db.NamedExec(query, payment)
	return err
}

func (r *PaymentRepository) UpdatePaymentStatus(statusId int, paymentId uuid.UUID) error {
	query := `UPDATE payments SET status_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, statusId, paymentId)
	return err
}
