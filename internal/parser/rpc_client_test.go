package parser

import (
	"testing"
)

func TestCallRPC(t *testing.T) {
	result, err := CallRPC("eth_blockNumber", []interface{}{})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result == "" {
		t.Errorf("expected a valid block number, got an empty string")
	}
}
