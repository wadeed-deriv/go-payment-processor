package application

import (
	"context"
	"log"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentGateway interface {
	Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail)
	Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail)
}

type PaymentRepository interface {
	GetClient(ctx context.Context, payment *entities.PaymentDetail) (*entities.Client, error)
}

type Paymentservice struct {
	payment PaymentRepository
}

func NewPaymentSerice(payment PaymentRepository) *Paymentservice {
	return &Paymentservice{payment: payment}
}

func (s *Paymentservice) MakeDeposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {

	var client *entities.Client
	client, err := s.payment.GetClient(ctx, paymentdetail)

	if err != nil {
		return err
	}

	var gateway = IdentifyPaymentGateway(client.Gateway)

	log.Println("error", err)

	log.Println("client", client)
	log.Println(paymentdetail)
	gateway.Deposit(ctx, paymentdetail)
	return nil
}

func (s *Paymentservice) MakeWithdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	var gateway = IdentifyPaymentGateway("B")
	log.Println(paymentdetail)
	gateway.Withdrawal(ctx, paymentdetail)
	return nil
}
