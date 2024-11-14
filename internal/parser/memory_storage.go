package parser

import "sync"

type MemoryStorage struct {
	addresses    map[string]bool
	transactions map[string][]Transaction
	currentBlock int
	mu           sync.RWMutex
}

// NewMemoryStorage creates a new instance of MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		addresses:    make(map[string]bool),
		transactions: make(map[string][]Transaction),
		currentBlock: 0,
	}
}
