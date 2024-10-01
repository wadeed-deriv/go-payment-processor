package paymentgatewayB

import (
	"context"
	"log"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (r *PaymentGateway) Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) {
	log.Println("Depositing to payment gateway B")
}

func (r *PaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) {
	//make http depoist call to gateway
}
