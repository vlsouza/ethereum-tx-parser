package parser

import (
	"testing"
)

func TestSubscribe(t *testing.T) {
	ethParser := NewEthereumParser()

	address := "0x12345"
	success := ethParser.Subscribe(address)
	if !success {
		t.Errorf("expected subscription to be successful for address %s", address)
	}

	// Test re-subscription
	success = ethParser.Subscribe(address)
	if success {
		t.Errorf("expected subscription to fail for already subscribed address %s", address)
	}
}

func TestGetCurrentBlock(t *testing.T) {
	ethParser := NewEthereumParser()

	block := ethParser.GetCurrentBlock()
	if block != 0 {
		t.Errorf("expected current block to be 0, got %d", block)
	}
}
