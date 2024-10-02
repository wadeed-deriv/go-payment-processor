package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

func TestNewPaymentRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPaymentRepository(db)
	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestGetClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPaymentRepository(db)
	clientID := "123"
	client := &entities.Client{
		ID:      1,
		Name:    "Test Client",
		Gateway: "Test Gateway",
		Balance: 100.0,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "gateway", "balance"}).
		AddRow(client.ID, client.Name, client.Gateway, client.Balance)

	mock.ExpectQuery("SELECT id, name, gateway, balance FROM client WHERE id = \\$1").
		WithArgs(clientID).
		WillReturnRows(rows)

	result, err := repo.GetClient(context.Background(), clientID)
	assert.NoError(t, err)
	assert.Equal(t, client, result)
}

func TestGetClient_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPaymentRepository(db)
	clientID := "123"

	mock.ExpectQuery("SELECT id, name, gateway, balance FROM client WHERE id = \\$1").
		WithArgs(clientID).
		WillReturnError(sql.ErrNoRows)

	result, err := repo.GetClient(context.Background(), clientID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "client not found", err.Error())
}

func TestUpdateClientBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPaymentRepository(db)
	client := &entities.Client{
		ID:      1,
		Balance: 200.0,
	}

	mock.ExpectExec("UPDATE client SET balance = \\$1 WHERE id = \\$2").
		WithArgs(client.Balance, client.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateClientBalance(context.Background(), client)
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewPaymentRepository(db)
	transaction := &entities.Transaction{
		ClientID: 1,
		Amount:   50.0,
		Type:     "credit",
	}

	mock.ExpectExec("INSERT INTO transaction \\(client_id, amount, type\\) VALUES \\(\\$1, \\$2, \\$3\\)").
		WithArgs(transaction.ClientID, transaction.Amount, transaction.Type).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateTransaction(context.Background(), transaction)
	assert.NoError(t, err)
}
