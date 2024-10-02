package paymentgatewayA

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
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
	client *http.Client
}

func NewPaymentGateway(client *http.Client) *PaymentGateway {
	return &PaymentGateway{
		client: client,
	}
}

func (r *PaymentGateway) Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	depositURL := os.Getenv("GATEWAY_A_URL")
	if depositURL == "" {
		depositURL = "http://127.0.0.1:3000/json/deposit"
	} else {
		depositURL = depositURL + "/deposit"
	}

	depositReq := Request{
		Amount:   paymentdetail.Amount,
		ClientID: paymentdetail.ID,
	}

	log.Println("deposit request : ", depositReq)

	err := r.sendRequest(ctx, depositReq, depositURL)
	if err != nil {
		log.Println("deposit failed due to error : ", err)
		return errors.New("deposit failed")
	}

	return nil
}

func (r *PaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	withdrawalURL := os.Getenv("GATEWAY_A_URL")
	if withdrawalURL == "" {
		withdrawalURL = "http://127.0.0.1:3000/json/withdrawal"
	} else {
		withdrawalURL = withdrawalURL + "/withdrawal"
	}

	withdrawalReq := Request{
		Amount:   paymentdetail.Amount,
		ClientID: paymentdetail.ID,
	}

	log.Println("withdrawal request : ", withdrawalReq)

	err := r.sendRequest(ctx, withdrawalReq, withdrawalURL)
	if err != nil {
		log.Println("withdrawal failed due to error : ", err)
		return errors.New("withdrawal failed")
	}

	return nil
}

func (r *PaymentGateway) sendRequest(ctx context.Context, request Request, url string) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return errors.New("gateway error")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("gateway error")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
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
