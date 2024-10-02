package main

import (
	"log"
	"os"

	"net/http"

	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/httphandler"
	"github.com/wadeed-deriv/go-payment-processor/internal/application"
	"github.com/wadeed-deriv/go-payment-processor/internal/db/postgres"
)

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
	httpClient := &http.Client{}
	paymentSerive := application.NewPaymentSerice(paymentRepo, httpClient)
	payment := httphandler.NewPaymentHandler(paymentSerive)
	server := httphandler.NewServer(payment)

	server.Start(":8080")
}
