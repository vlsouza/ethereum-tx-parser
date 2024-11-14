package parser

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const rpcURL = "https://ethereum-rpc.publicnode.com"

// RPCRequest represents the structure of a JSON-RPC request
type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// RPCResponse represents the structure of a JSON-RPC response
type RPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

// CallRPC sends a JSON-RPC request and returns the result
func CallRPC(method string, params []interface{}) (string, error) {
	request := RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var rpcResponse RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResponse); err != nil {
		return "", err
	}
	return rpcResponse.Result, nil
}
