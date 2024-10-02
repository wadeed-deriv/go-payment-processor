package httphandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/httphandler"
	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

// Mocking the PaymentService
type MockPaymentService struct {
	mock.Mock
}

func (m *MockPaymentService) MakeDeposit(ctx context.Context, paymentDetail *entities.PaymentDetail) error {
	args := m.Called(ctx, paymentDetail)
	return args.Error(0)
}

func (m *MockPaymentService) MakeWithdrawal(ctx context.Context, paymentDetail *entities.PaymentDetail) error {
	args := m.Called(ctx, paymentDetail)
	return args.Error(0)
}

func (m *MockPaymentService) TransactionUpdate(ctx context.Context, transactionUpdate *entities.TransactionUpdate) error {
	args := m.Called(ctx, transactionUpdate)
	return args.Error(0)
}

func TestMakeDeposit_Success(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	paymentDetail := entities.PaymentDetail{ID: "client123", Amount: 100.0}
	mockService.On("MakeDeposit", mock.Anything, &paymentDetail).Return(nil)

	reqBody, _ := json.Marshal(paymentDetail)
	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.MakeDeposit(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "Deposit made successfully", response["message"])
	mockService.AssertExpectations(t)
}

func TestMakeDeposit_Failure(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	paymentDetail := entities.PaymentDetail{ID: "client123", Amount: 100.0}
	mockService.On("MakeDeposit", mock.Anything, &paymentDetail).Return(errors.New("deposit failed"))

	reqBody, _ := json.Marshal(paymentDetail)
	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.MakeDeposit(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "failed", response["status"])
	assert.Equal(t, "deposit failed", response["message"])
	mockService.AssertExpectations(t)
}

func TestMakeWithdrawal_Success(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	paymentDetail := entities.PaymentDetail{ID: "client123", Amount: 50.0}
	mockService.On("MakeWithdrawal", mock.Anything, &paymentDetail).Return(nil)

	reqBody, _ := json.Marshal(paymentDetail)
	req := httptest.NewRequest(http.MethodPost, "/withdrawal", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.MakeWithdrawal(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "Withdrawal made successfully", response["message"])
	mockService.AssertExpectations(t)
}

func TestMakeWithdrawal_Failure(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	paymentDetail := entities.PaymentDetail{ID: "client123", Amount: 50.0}
	mockService.On("MakeWithdrawal", mock.Anything, &paymentDetail).Return(errors.New("withdrawal failed"))

	reqBody, _ := json.Marshal(paymentDetail)
	req := httptest.NewRequest(http.MethodPost, "/withdrawal", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.MakeWithdrawal(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "failed", response["status"])
	assert.Equal(t, "withdrawal failed", response["message"])
	mockService.AssertExpectations(t)
}

func TestTransactionUpdate_Success(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	transactionUpdate := entities.TransactionUpdate{AccountID: "client123", Amount: 200.0, TransactionType: "DEPOSIT"}
	mockService.On("TransactionUpdate", mock.Anything, &transactionUpdate).Return(nil)

	reqBody, _ := json.Marshal(transactionUpdate)
	req := httptest.NewRequest(http.MethodPost, "/transaction/update", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.TransactionUpdate(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "Transaction updated successfully", response["message"])
	mockService.AssertExpectations(t)
}

func TestTransactionUpdate_Failure(t *testing.T) {
	mockService := new(MockPaymentService)
	handler := httphandler.NewPaymentHandler(mockService)

	transactionUpdate := entities.TransactionUpdate{AccountID: "client123", Amount: 200.0, TransactionType: "DEPOSIT"}
	mockService.On("TransactionUpdate", mock.Anything, &transactionUpdate).Return(errors.New("transaction update failed"))

	reqBody, _ := json.Marshal(transactionUpdate)
	req := httptest.NewRequest(http.MethodPost, "/transaction/update", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	handler.TransactionUpdate(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var response map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "failed", response["status"])
	assert.Equal(t, "transaction update failed", response["message"])
	mockService.AssertExpectations(t)
}
