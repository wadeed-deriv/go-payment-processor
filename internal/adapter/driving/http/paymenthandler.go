package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wadeed-deriv/go-payment-processor/internal/application"
	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentHandler struct {
	service *application.Paymentservice
}

func NewPaymentHandler(service *application.Paymentservice) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) MakeDeposit(w http.ResponseWriter, r *http.Request) {
	var paymentDetail entities.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&paymentDetail); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Println(paymentDetail)

	if err := h.service.MakeDeposit(r.Context(), &paymentDetail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PaymentHandler) MakeWithdrawal(w http.ResponseWriter, r *http.Request) {
	var paymentDetail entities.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&paymentDetail); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.service.MakeWithdrawal(r.Context(), &paymentDetail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
