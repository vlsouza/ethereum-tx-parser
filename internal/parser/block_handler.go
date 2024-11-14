package parser

import (
	"fmt"
	"log"
	"strconv"
)

// UpdateCurrentBlock fetches the latest block number and updates storage
func (p *EthereumParser) UpdateCurrentBlock() error {
	result, err := CallRPC("eth_blockNumber", []interface{}{})
	if err != nil {
		return err
	}

	blockNumber, err := strconv.ParseInt(result, 0, 64)
	if err != nil {
		return fmt.Errorf("failed to parse block number: %v", err)
	}

	p.storage.mu.Lock()
	defer p.storage.mu.Unlock()
	p.storage.currentBlock = int(blockNumber)
	log.Printf("Updated current block to %d", p.storage.currentBlock)
	return nil
}
