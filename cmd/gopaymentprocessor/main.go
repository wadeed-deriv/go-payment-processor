package main

import (
	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driving/http"
	"github.com/wadeed-deriv/go-payment-processor/internal/application"
)

func main() {
	var paymentRepo application.PaymentGateway

	paymentSerive := application.NewPaymentSerice(paymentRepo)
	payment := http.NewPaymentHandler(paymentSerive)
	server := http.NewServer(payment)

	server.Start(":8080")
}
