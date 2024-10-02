package httphandler

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

// setup endpoints
func NewServer(handler *PaymentHandler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/deposit", handler.MakeDeposit)
	mux.HandleFunc("/withdrawal", handler.MakeWithdrawal)
	mux.HandleFunc("/update", handler.TransactionUpdate)
	return &Server{mux: mux}
}

// Start the server
func (s *Server) Start(addr string) {
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}
