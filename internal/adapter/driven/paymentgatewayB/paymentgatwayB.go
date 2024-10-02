package paymentgatewayB

import (
	"bytes"
	"context"
	"encoding/xml"
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

	depositURL := os.Getenv("GATEWAY_B_URL")
	if depositURL == "" {
		depositURL = "http://127.0.0.1:3000/xml/deposit"
	} else {
		depositURL = depositURL + "/deposit"
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

	withdrawalURL := os.Getenv("GATEWAY_B_URL")
	if withdrawalURL == "" {
		withdrawalURL = "http://127.0.0.1:3000/xml/withdrawal"
	} else {
		withdrawalURL = withdrawalURL + "/withdrawal"
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
	reqBody, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", depositURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/xml")

	log.Println("Sending request to gateway", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("request failed with status: " + resp.Status)
	}

	return nil
}
