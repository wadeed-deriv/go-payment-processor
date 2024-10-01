package application

import (
	"context"
	"errors"
	"log"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentGateway interface {
	Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error
	Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error
}

type PaymentRepository interface {
	GetClient(ctx context.Context, payment *entities.PaymentDetail) (*entities.Client, error)
	UpdateClientBalance(ctx context.Context, client *entities.Client) error
	CreateTransaction(ctx context.Context, transaction *entities.Transaction) error
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
		return errors.New("client not found")
	}
	var gateway = IdentifyPaymentGateway(client.Gateway)

	log.Println(paymentdetail)
	err = gateway.Deposit(ctx, paymentdetail)

	if err != nil {
		return errors.New("deposit failed")
	}
	log.Println("Deposited")

	client.Balance += paymentdetail.Amount
	err = s.payment.UpdateClientBalance(ctx, client)
	if err != nil {
		return errors.New("deposit failed")
	}

	transaction := &entities.Transaction{
		ClientID: client.ID,
		Amount:   paymentdetail.Amount,
		Type:     "DEPOSIT",
	}
	s.payment.CreateTransaction(ctx, transaction)
	return nil
}

func (s *Paymentservice) MakeWithdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	var client *entities.Client
	client, err := s.payment.GetClient(ctx, paymentdetail)

	if err != nil {
		return errors.New("client not found")
	}
	var gateway = IdentifyPaymentGateway(client.Gateway)

	log.Println(paymentdetail)
	err = gateway.Withdrawal(ctx, paymentdetail)

	if err != nil {
		return errors.New("withdrawal failed")
	}
	log.Println("Withdrawn")

	client.Balance -= paymentdetail.Amount
	err = s.payment.UpdateClientBalance(ctx, client)
	if err != nil {
		return errors.New("withdrawal failed")
	}

	transaction := &entities.Transaction{
		ClientID: client.ID,
		Amount:   paymentdetail.Amount,
		Type:     "WITHDRAWAL",
	}
	s.payment.CreateTransaction(ctx, transaction)
	return nil
}
