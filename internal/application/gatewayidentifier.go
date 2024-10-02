package application

import (
	"log"

	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driven/paymentgatewayA"
	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driven/paymentgatewayB"
)

// Base on client registration, identify the payment gateway
func IdentifyPaymentGateway(gateway string) PaymentGateway {

	switch gateway {
	case "A":
		return paymentgatewayA.NewPaymentGateway()
	case "B":
		return paymentgatewayB.NewPaymentGateway()
	default:
		log.Println("Gateway not found")
		return nil
	}
}
