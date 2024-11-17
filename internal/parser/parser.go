package parser

import (
	"ethereum-tx-parser/internal/rpcclient"
	"log"
)

// Transaction struct represents a simple transaction model
type Transaction struct {
	From  string
	To    string
	Value string
}

// Parser interface defines the main methods
type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}

// EthereumParser struct implements the Parser interface
type EthereumParser struct {
	rpcClient rpcclient.Client
	storage   *MemoryStorage
}

// NewEthereumParser initializes a new EthereumParser
func NewEthereumParser(c rpcclient.Client) *EthereumParser {
	return &EthereumParser{
		rpcClient: c,
		storage:   NewMemoryStorage(),
	}
}

// GetCurrentBlock returns the last parsed block number
func (p *EthereumParser) GetCurrentBlock() int {
	return p.storage.currentBlock
}

// Subscribe adds an address to the observer list
func (p *EthereumParser) Subscribe(address string) bool {
	p.storage.mu.Lock()
	defer p.storage.mu.Unlock()
	if _, exists := p.storage.addresses[address]; !exists {
		p.storage.addresses[address] = true
		log.Printf("Address %s subscribed successfully", address)
	}
	log.Printf("Address %s is already subscribed", address)
	return true
}

// GetTransactions returns a list of transactions for a given address
func (p *EthereumParser) GetTransactions(address string) []Transaction {
	p.storage.mu.RLock()
	defer p.storage.mu.RUnlock()
	return p.storage.transactions[address]
}
