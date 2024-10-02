package main

import (
	"log"
	"os"

	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/http"
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
	paymentSerive := application.NewPaymentSerice(paymentRepo)
	payment := http.NewPaymentHandler(paymentSerive)
	server := http.NewServer(payment)

	server.Start(":8080")
}
