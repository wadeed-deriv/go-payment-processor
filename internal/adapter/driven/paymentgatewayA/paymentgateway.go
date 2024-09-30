package paymentgatewayA

import (
	"context"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (r *PaymentGateway) deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) {
	//make http depoist call to gateway
}

func (r *PaymentGateway) Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) {
	//make http depoist call to gateway
}
