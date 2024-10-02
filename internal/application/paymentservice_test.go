package application

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

// MockPaymentRepository is a mock implementation of the PaymentRepository interface
type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) GetClient(ctx context.Context, clientid string) (*entities.Client, error) {
	args := m.Called(ctx, clientid)
	return args.Get(0).(*entities.Client), args.Error(1)
}

func (m *MockPaymentRepository) UpdateClientBalance(ctx context.Context, client *entities.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateTransaction(ctx context.Context, transaction *entities.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

// MockPaymentGateway is a mock implementation of the PaymentGateway interface
type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	args := m.Called(ctx, paymentdetail)
	return args.Error(0)
}

func (m *MockPaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	args := m.Called(ctx, paymentdetail)
	return args.Error(0)
}

// IdentifyPaymentGateway is a mock function to return a mock payment gateway
func MockIdentifyPaymentGateway(gateway string, httpClient *http.Client) PaymentGateway {
	return &MockPaymentGateway{}
}

func TestMakeDeposit(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	mockGateway := new(MockPaymentGateway)
	httpClient := &http.Client{}
	service := NewPaymentSerice(mockRepo, httpClient)

	client := &entities.Client{
		ID:      1,
		Balance: 100,
		Gateway: "A",
	}
	paymentDetail := &entities.PaymentDetail{
		ID:     "1",
		Amount: 50,
	}

	mockRepo.On("GetClient", mock.Anything, "1").Return(client, nil)
	mockGateway.On("Deposit", mock.Anything, paymentDetail).Return(nil)
	mockRepo.On("UpdateClientBalance", mock.Anything, client).Return(nil)
	mockRepo.On("CreateTransaction", mock.Anything, mock.Anything).Return(nil)

	err := service.MakeDeposit(context.Background(), paymentDetail)
	assert.NoError(t, err)
	assert.Equal(t, float64(150), client.Balance)
	mockRepo.AssertExpectations(t)
}

func TestMakeWithdrawal(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	mockGateway := new(MockPaymentGateway)
	httpClient := &http.Client{}
	service := NewPaymentSerice(mockRepo, httpClient)

	client := &entities.Client{
		ID:      1,
		Balance: 100,
		Gateway: "A",
	}
	paymentDetail := &entities.PaymentDetail{
		ID:     "1",
		Amount: 50,
	}

	mockRepo.On("GetClient", mock.Anything, "1").Return(client, nil)
	mockGateway.On("Withdrawal", mock.Anything, paymentDetail).Return(nil)
	mockRepo.On("UpdateClientBalance", mock.Anything, client).Return(nil)
	mockRepo.On("CreateTransaction", mock.Anything, mock.Anything).Return(nil)

	err := service.MakeWithdrawal(context.Background(), paymentDetail)
	assert.NoError(t, err)
	assert.Equal(t, float64(50), client.Balance)
	mockRepo.AssertExpectations(t)
}

func TestTransactionUpdate(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	httpClient := &http.Client{}
	service := NewPaymentSerice(mockRepo, httpClient)

	client := &entities.Client{
		ID:      1,
		Balance: 100,
	}
	transactionUpdate := &entities.TransactionUpdate{
		AccountID:       "1",
		Amount:          50,
		TransactionType: "DEPOSIT",
	}

	mockRepo.On("GetClient", mock.Anything, "1").Return(client, nil)
	mockRepo.On("UpdateClientBalance", mock.Anything, client).Return(nil)
	mockRepo.On("CreateTransaction", mock.Anything, mock.Anything).Return(nil)

	err := service.TransactionUpdate(context.Background(), transactionUpdate)
	assert.NoError(t, err)
	assert.Equal(t, float64(150), client.Balance)
	mockRepo.AssertExpectations(t)
}
