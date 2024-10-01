package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) GetClient(ctx context.Context, payment *entities.PaymentDetail) (*entities.Client, error) {
	client := &entities.Client{}
	err := r.db.QueryRowContext(ctx, "SELECT id, name, gateway, balance FROM client WHERE id = $1", payment.ID).
		Scan(&client.ID, &client.Name, &client.Gateway, &client.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle no rows returned (e.g. client not found)
			return nil, errors.New("client not found")
		}
		return nil, err
	}

	return client, nil
}

// func (r *PaymentRepository) GetClient(ctx context.Context, payment *entities.PaymentDetail) error {
// 	_, err := r.db.ExecContext(ctx, "SELECT * FROM client where id =  $1", payment.ID)
// 	return err
// }
