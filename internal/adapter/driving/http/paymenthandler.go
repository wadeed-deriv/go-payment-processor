package http

import (
	"encoding/json"
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
		h.respond(w, http.StatusBadRequest, "failed", "Invalid request payload")
		return
	}

	if err := h.service.MakeDeposit(r.Context(), &paymentDetail); err != nil {
		h.respond(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	h.respond(w, http.StatusOK, "success", "Deposit made successfully")
}

func (h *PaymentHandler) MakeWithdrawal(w http.ResponseWriter, r *http.Request) {
	var paymentDetail entities.PaymentDetail
	if err := json.NewDecoder(r.Body).Decode(&paymentDetail); err != nil {
		h.respond(w, http.StatusBadRequest, "failed", "Invalid request payload")
		return
	}

	if err := h.service.MakeWithdrawal(r.Context(), &paymentDetail); err != nil {
		h.respond(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	h.respond(w, http.StatusOK, "success", "Withdrawal made successfully")
}

func (h *PaymentHandler) respond(w http.ResponseWriter, statusCode int, status string, message string) {
	w.WriteHeader(statusCode)
	response := map[string]string{"status": status, "message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
