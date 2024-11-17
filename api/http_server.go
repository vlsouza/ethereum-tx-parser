package api

import (
	"log"
	"net/http"

	"encoding/json"
	_ "ethereum-tx-parser/docs"
	"ethereum-tx-parser/internal/parser"
	httpSwagger "github.com/swaggo/http-swagger"
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

	// Add the Swagger UI handler
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Printf("Server started at port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// @Summary Get current block
// @Description Get the latest parsed block
// @Produce json
// @Success 200 {object} map[string]int "A map where the key is 'currentBlock' and the value is the block number"
// @Router /currentBlock [get]
func (s *Server) getCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

// @Summary Subscribe given an address
// @Description Subscribe to notifications for incoming/outgoing transactions for a specific Ethereum address
// @Produce json
// @Param address query string true "Ethereum address to subscribe to"
// @Success 200 {object} map[string]bool "A map where the key is 'subscribed' and the value is a boolean indicating success"
// @Failure 400 {string} string "Address is required"
// @Failure 409 {string} string "Already subscribed"
// @Router /subscribe [get]
func (s *Server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	success := s.parser.Subscribe(address)
	if !success {
		http.Error(w, "already subscribed", http.StatusConflict)
	}
	json.NewEncoder(w).Encode(map[string]bool{"subscribed": success})
}

// @Summary Get transactions given an address
// @Description Retrieve inbound and outbound transactions for a subscribed Ethereum address
// @Produce json
// @Param address query string true "Ethereum address to retrieve transactions for"
// @Success 200 {array} parser.Transaction
// @Failure 400 {string} string "Address is required"
// @Router /transactions [get]
func (s *Server) getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	transactions := s.parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}
