package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	_ "ethereum-tx-parser/docs"
	"ethereum-tx-parser/internal/parser"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/block/current", s.getCurrentBlockHandler).Methods("GET")
	router.HandleFunc("/subscribe/{address}", s.subscribeHandler).Methods("POST")
	router.HandleFunc("/transactions/{address}", s.getTransactionsHandler).Methods("GET")

	// Add the Swagger UI handler
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Printf("Server started at port %s\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// @Summary Get current block
// @Description Get the latest parsed block
// @Produce json
// @Success 200 {object} map[string]int "A map where the key is 'currentBlock' and the value is the block number"
// @Router /block/current [get]
func (s *Server) getCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

// @Summary Subscribe given an address
// @Description Subscribe to notifications for incoming/outgoing transactions for a specific Ethereum address
// @Produce json
// @Param address path string true "Ethereum address to subscribe to"
// @Success 200 {object} map[string]bool "A map where the key is 'subscribed' and the value is a boolean indicating success"
// @Failure 400 {string} string "Address is required"
// @Failure 409 {string} string "Already subscribed"
// @Router /subscribe/{address} [post]
func (s *Server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	address, err := getAddressFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	success := s.parser.Subscribe(address)
	if !success {
		http.Error(w, "already subscribed", http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"subscribed": success}) //can return more info if needed
}

// @Summary Get transactions given an address
// @Description Retrieve inbound and outbound transactions for a subscribed Ethereum address
// @Produce json
// @Param address path string true "Ethereum address to retrieve transactions for"
// @Success 200 {array} parser.Transaction
// @Failure 400 {string} string "Address is required"
// @Router /transactions/{address} [get]
func (s *Server) getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address, err := getAddressFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transactions := s.parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}

// Helper function to get and validate the address from the request
func getAddressFromRequest(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	address := vars["address"]
	if address == "" {
		return "", errors.New("address is required")
	}
	return address, nil
}
