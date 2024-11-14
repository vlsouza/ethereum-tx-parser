package api

import (
	"encoding/json"
	"ethereum-tx-parser/internal/parser"
	"log"
	"net/http"
)

type Server struct {
	parser *parser.EthereumParser
}

// NewServer initializes a new HTTP server
func NewServer(p *parser.EthereumParser) *Server {
	return &Server{parser: p}
}

// StartHTTPServer starts the HTTP server on a specified port
func (s *Server) StartHTTPServer(port string) {
	http.HandleFunc("/currentBlock", s.getCurrentBlockHandler)
	http.HandleFunc("/subscribe", s.subscribeHandler)
	http.HandleFunc("/transactions", s.getTransactionsHandler)

	log.Printf("Server started at port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func (s *Server) getCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

func (s *Server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	success := s.parser.Subscribe(address)
	json.NewEncoder(w).Encode(map[string]bool{"subscribed": success})
}

func (s *Server) getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	transactions := s.parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}
