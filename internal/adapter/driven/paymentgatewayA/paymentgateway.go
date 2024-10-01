package paymentgatewayA

import (
	"context"
	"errors"
	"os"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"

	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Amount   float64 `json:"amount"`
	ClientID string  `json:"clientID"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (r *PaymentGateway) Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {

	depositURL := os.Getenv("GATEWAY_A_URL") + "/deposit"
	if depositURL == "" {
		depositURL = "http://127.0.0.1:3000/json/deposit"
	}

	depositReq := Request{
		Amount:   paymentdetail.Amount,
		ClientID: paymentdetail.ID,
	}

	err := sendRequest(ctx, depositReq, depositURL)
	if err != nil {
		return errors.New("deposit failed")
	}

	return nil
}

func (r *PaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {

	withdrawalURL := os.Getenv("GATEWAY_A_URL") + "/withdrawal"
	if withdrawalURL == "" {
		withdrawalURL = "http://127.0.0.1:3000/json/withdrawal"
	}

	withdrawalReq := Request{
		Amount:   paymentdetail.Amount,
		ClientID: paymentdetail.ID,
	}

	err := sendRequest(ctx, withdrawalReq, withdrawalURL)
	if err != nil {
		return errors.New("withdrawal failed")
	}

	return nil
}

func sendRequest(ctx context.Context, request Request, depositURL string) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return errors.New("gateway error")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", depositURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("gateway error")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("gateway error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("gateway error")
	}

	var depositResp Response
	if err := json.NewDecoder(resp.Body).Decode(&depositResp); err != nil {
		return errors.New("gateway error")
	}
	return nil
}
