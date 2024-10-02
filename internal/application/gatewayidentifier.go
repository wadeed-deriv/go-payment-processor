package application

import (
	"log"
	"net/http"

	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driven/paymentgatewayA"
	"github.com/wadeed-deriv/go-payment-processor/internal/adapter/driven/paymentgatewayB"
)

// Base on client registration, identify the payment gateway
func IdentifyPaymentGateway(gateway string, httpClient *http.Client) PaymentGateway {

	switch gateway {
	case "A":
		return paymentgatewayA.NewPaymentGateway(httpClient)
	case "B":
		return paymentgatewayB.NewPaymentGateway(httpClient)
	default:
		log.Println("Gateway not found")
		return nil
	}
}
