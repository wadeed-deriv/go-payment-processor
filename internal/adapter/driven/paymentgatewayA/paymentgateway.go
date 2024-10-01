package paymentgatewayA

import (
	"context"
	"log"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"

	"bytes"
	"encoding/json"
	"net/http"
)

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (r *PaymentGateway) Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) {

	type DepositRequest struct {
		Amount   float64 `json:"amount"`
		ClientID string  `json:"clientID"`
	}

	type DepositResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	depositURL := "https://thirdpartyservice.com/api/deposit"

	depositReq := DepositRequest{
		Amount:   paymentdetail.Amount,
		ClientID: paymentdetail.ID,
	}

	jsonData, err := json.Marshal(depositReq)
	if err != nil {
		log.Printf("Error marshalling request: %v", err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", depositURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK response: %v", resp.Status)
		return
	}

	var depositResp DepositResponse
	if err := json.NewDecoder(resp.Body).Decode(&depositResp); err != nil {
		log.Printf("Error decoding response: %v", err)
		return
	}

	log.Printf("Deposit response: %v", depositResp)
}

func (r *PaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) {
	//make http depoist call to gateway
}
