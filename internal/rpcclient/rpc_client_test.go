package rpcclient

import (
	"testing"
)

func TestCall(t *testing.T) {
	rpcClient := NewMockedCallSuccess()
	result, err := rpcClient.Call("eth_blockNumber", []interface{}{})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result == "" {
		t.Errorf("expected a valid block number, got an empty string")
	}
}
