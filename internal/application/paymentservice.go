package application

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

// PaymentGateway interface
type PaymentGateway interface {
	Deposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error
	Withdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error
}

// Payment DB repository interface
type PaymentRepository interface {
	GetClient(ctx context.Context, clientid string) (*entities.Client, error)
	UpdateClientBalance(ctx context.Context, client *entities.Client) error
	CreateTransaction(ctx context.Context, transaction *entities.Transaction) error
}

type Paymentservice struct {
	payment    PaymentRepository
	httpClient *http.Client
}

func NewPaymentSerice(payment PaymentRepository, httpClient *http.Client) *Paymentservice {
	return &Paymentservice{payment: payment, httpClient: httpClient}
}

/**
 * MakeDeposit
 * @summary Process the clients request to make a deposit
 * @param ctx context.Context
 * @param paymentdetail *entities.PaymentDetail
 * @return error
 */
func (s *Paymentservice) MakeDeposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error {

	var client *entities.Client
	client, err := s.payment.GetClient(ctx, paymentdetail.ID)

	if err != nil {
		return errors.New("client not found")
	}
	var gateway = IdentifyPaymentGateway(client.Gateway, s.httpClient)

	log.Println(paymentdetail)
	err = gateway.Deposit(ctx, paymentdetail)

	transaction := &entities.Transaction{
		ClientID: client.ID,
		Amount:   paymentdetail.Amount,
		Type:     "DEPOSIT",
	}

	if err == nil {
		client.Balance += paymentdetail.Amount
		err = s.payment.UpdateClientBalance(ctx, client)

		if err == nil {
			log.Println("Deposited")
			transaction.Status = "COMPLETED"
			log.Println(transaction)
			err = s.payment.CreateTransaction(ctx, transaction)
			if err != nil {
				log.Println(err)
			}
			return nil
		} else {
			transaction.Status = "FAILED"
		}
	} else {
		transaction.Status = "FAILED"
	}

	s.payment.CreateTransaction(ctx, transaction)
	return errors.New("deposit failed")
}

/**
* MakeWithdrawal
* @summary Process the clients request to make a withdrawal
* @param ctx context.Context
* @param paymentdetail *entities.PaymentDetail
* @return error
 */
func (s *Paymentservice) MakeWithdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error {
	var client *entities.Client
	client, err := s.payment.GetClient(ctx, paymentdetail.ID)

	if err != nil {
		return errors.New("client not found")
	}
	var gateway = IdentifyPaymentGateway(client.Gateway, s.httpClient)

	log.Println(paymentdetail)
	err = gateway.Withdrawal(ctx, paymentdetail)

	transaction := &entities.Transaction{
		ClientID: client.ID,
		Amount:   paymentdetail.Amount,
		Type:     "WITHDRAWAL",
	}

	if err == nil {
		client.Balance -= paymentdetail.Amount
		err = s.payment.UpdateClientBalance(ctx, client)

		if err == nil {
			log.Println("Withdrawn")
			transaction.Status = "COMPLETED"
			s.payment.CreateTransaction(ctx, transaction)
			return nil
		} else {
			transaction.Status = "FAILED"
		}
	} else {
		transaction.Status = "FAILED"
	}
	s.payment.CreateTransaction(ctx, transaction)
	return errors.New("withdrawal failed")
}

/**
 * TransactionUpdate
 * @summary Update the transaction
 * @param ctx context.Context
 * @param transactionUpdate *entities.TransactionUpdate
 * @return error
 */
func (s *Paymentservice) TransactionUpdate(ctx context.Context, transactionUpdate *entities.TransactionUpdate) error {
	var client *entities.Client
	client, err := s.payment.GetClient(ctx, transactionUpdate.AccountID)

	if err != nil {
		return errors.New("client not found")
	}

	if transactionUpdate.TransactionType == "DEPOSIT" || transactionUpdate.TransactionType == "WITHDRAWAL_REVERSAL" {
		client.Balance += transactionUpdate.Amount
	} else if transactionUpdate.TransactionType == "WITHDRAWAL" || transactionUpdate.TransactionType == "DEPOSIT_REVERSAL" {
		client.Balance -= transactionUpdate.Amount
	}

	transaction := &entities.Transaction{
		ClientID: client.ID,
		Amount:   transactionUpdate.Amount,
		Type:     transactionUpdate.TransactionType,
	}

	err = s.payment.UpdateClientBalance(ctx, client)
	if err == nil {
		log.Println("Transaction Updated")
		transaction.Status = "COMPLETED"
		s.payment.CreateTransaction(ctx, transaction)
		return nil
	}
	transaction.Status = "FAILED"
	s.payment.CreateTransaction(ctx, transaction)
	return errors.New("transaction update failed")
}
