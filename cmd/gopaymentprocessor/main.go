package main

import (
	"log"
	"os"

	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/http"
	"github.com/wadeed-deriv/go-payment-processor/internal/application"
	"github.com/wadeed-deriv/go-payment-processor/internal/db/postgres"
)

func main() {

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

	//return payment gateway based on the type
	paymentSerive := application.NewPaymentSerice(paymentRepo)
	payment := http.NewPaymentHandler(paymentSerive)
	server := http.NewServer(payment)

	server.Start(":8080")
}
