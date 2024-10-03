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

func (r *PaymentRepository) GetClient(ctx context.Context, clientid string) (*entities.Client, error) {
	client := &entities.Client{}
	err := r.db.QueryRowContext(ctx, "SELECT id, name, gateway, balance FROM client WHERE id = $1", clientid).
		Scan(&client.ID, &client.Name, &client.Gateway, &client.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}
	return client, nil
}

func (r *PaymentRepository) UpdateClientBalance(ctx context.Context, client *entities.Client) error {
	_, err := r.db.ExecContext(ctx, "UPDATE client SET balance = $1 WHERE id = $2", client.Balance, client.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) CreateTransaction(ctx context.Context, transaction *entities.Transaction) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO transaction (client_id, amount, type, status) VALUES ($1, $2, $3, $4)",
		transaction.ClientID, transaction.Amount, transaction.Type, transaction.Status)
	if err != nil {
		return err
	}
	return nil
}
