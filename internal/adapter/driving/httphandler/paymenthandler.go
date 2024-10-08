package httphandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wadeed-deriv/go-payment-processor/internal/domain/entities"
)

type PaymentService interface {
	MakeDeposit(ctx context.Context, paymentdetail *entities.PaymentDetail) error
	MakeWithdrawal(ctx context.Context, paymentdetail *entities.PaymentDetail) error
	TransactionUpdate(ctx context.Context, transactionUpdate *entities.TransactionUpdate) error
}

type PaymentHandler struct {
	service PaymentService
}

func NewPaymentHandler(service PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

/**
 * @api {post} /deposit Make Deposit
 * @apiName MakeDeposit
 * @apiGroup Payment
 *
 * @apiParam {Number} amount Amount to deposit.
 * @apiParam {String} clientID client identifier.
 *
 * @apiSuccess {String} status Status of the request.
 * @apiSuccess {String} message Message of the request.
 *
 **/
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

/**
 * @api {post} /withdrawal Make Withdrawal
 * @apiName MakeWithdrawal
 * @apiGroup Payment
 *
 * @apiParam {Number} amount Amount to withdraw.
 * @apiParam {String} clientID client identifier.
 *
 * @apiSuccess {String} status Status of the request.
 * @apiSuccess {String} message Message of the request.
 *
 **/
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

/**
 * @api {post} /update Update Transaction
* @apiName TransactionUpdate
* @apiGroup Payment
*
* @apiParam {String} transactionID Transaction identifier.
* @apiParam {String} status Transaction status.
* @apiParam {String} message Transaction message.
*
* @apiSuccess {String} status Status of the request.
* @apiSuccess {String} message Message of the request.
 **/
func (h *PaymentHandler) TransactionUpdate(w http.ResponseWriter, r *http.Request) {
	var transactionUpdate entities.TransactionUpdate
	if err := json.NewDecoder(r.Body).Decode(&transactionUpdate); err != nil {
		h.respond(w, http.StatusBadRequest, "failed", "Invalid request payload")
		return
	}

	if err := h.service.TransactionUpdate(r.Context(), &transactionUpdate); err != nil {
		h.respond(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	h.respond(w, http.StatusOK, "success", "Transaction updated successfully")
}

func (h *PaymentHandler) respond(w http.ResponseWriter, statusCode int, status string, message string) {
	w.WriteHeader(statusCode)
	response := map[string]string{"status": status, "message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
