package parser

import (
	"strings"
	"testing"
	
	"ethereum-tx-parser/internal/rpcclient"
)

func TestUpdateCurrentBlock_Success(t *testing.T) {
	rpcClient := rpcclient.NewMockedCallSuccess()
	ethParser := NewEthereumParser(rpcClient)

	err := ethParser.UpdateCurrentBlock()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verifies if the block was updated as expected
	if ethParser.GetCurrentBlock() != 1207 {
		t.Errorf("expected current block to be 1207, got %d", ethParser.GetCurrentBlock())
	}
}

func TestUpdateCurrentBlock_RPCError(t *testing.T) {
	rpcClient := rpcclient.NewMockedCallError()
	ethParser := NewEthereumParser(rpcClient)

	err := ethParser.UpdateCurrentBlock()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestUpdateCurrentBlock_InvalidBlockNumber(t *testing.T) {
	rpcClient := rpcclient.NewMockedCallInvalidNumber()
	ethParser := NewEthereumParser(rpcClient)

	err := ethParser.UpdateCurrentBlock()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to parse block number") {
		t.Fatalf("expected error message to contain 'failed to parse block number', got: %v", err)
	}
}
