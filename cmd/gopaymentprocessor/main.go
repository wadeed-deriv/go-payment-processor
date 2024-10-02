package main

import (
	"errors"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/cenkalti/backoff/v4"
	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/httphandler"
	"github.com/wadeed-deriv/go-payment-processor/internal/application"
	"github.com/wadeed-deriv/go-payment-processor/internal/db/postgres"
)

func retryClient() *http.Client {
	return &http.Client{
		Transport: &retryTransport{
			Transport: http.DefaultTransport,
		},
	}
}

type retryTransport struct {
	Transport http.RoundTripper
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	// Create exponential backoff with max elapsed time of 2 minutes
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 2 * time.Minute

	err = backoff.Retry(func() error {
		resp, err = t.Transport.RoundTrip(req)
		if err != nil {
			log.Printf("Retrying request due to error: %v", err)
			return err
		}

		// Retry for any status code >= 500 (server error)
		if resp.StatusCode >= 500 {
			log.Printf("Retrying due to status code: %d", resp.StatusCode)
			return errors.New("server error")
		}

		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {

	/**
	* Initializing database repository
	**/
	var paymentRepo application.PaymentRepository
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		connStr = "postgres://user:password@localhost:5434/payment?sslmode=disable"
	}
	db, err := postgres.NewPostgresConnection(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()
	paymentRepo = postgres.NewPaymentRepository(db)

	// Initializing payment service and injecting dependencies
	httpClient := retryClient()
	paymentSerive := application.NewPaymentSerice(paymentRepo, httpClient)
	payment := httphandler.NewPaymentHandler(paymentSerive)
	server := httphandler.NewServer(payment)

	server.Start(":8080")
}
