package application

import (
	"context"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentGateway interface {
	Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error
	Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error
}

type Paymentservice struct {
	gateway PaymentGateway
}

func NewPaymentSerice(gateway PaymentGateway) *Paymentservice {
	return &Paymentservice{gateway: gateway}
}

func (s *Paymentservice) MakeDeposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	return s.gateway.Deposit(ctx, paymentdetail)
}

func (s *Paymentservice) MakeWithdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	return s.gateway.Withdrawal(ctx, paymentdetail)
}
