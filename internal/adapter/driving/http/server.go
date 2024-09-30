package http

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(handler *PaymentHandler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/deposit", handler.MakeDeposit)
	mux.HandleFunc("/withdrawal", handler.MakeWithdrawal)
	return &Server{mux: mux}
}

func (s *Server) Start(addr string) {
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}
