package rpcclient

import "errors"

type MockedClientSuccess struct {
}

func NewMockedCallSuccess() Client {
	return &MockedClientSuccess{}
}

type MockedClientError struct {
}

func NewMockedCallError() Client {
	return &MockedClientError{}
}

type MockedClientInvalidNumber struct {
}

func NewMockedCallInvalidNumber() Client {
	return &MockedClientInvalidNumber{}
}

// Call..
func (m MockedClientSuccess) Call(method string, params []interface{}) (string, error) {
	if method == "eth_blockNumber" {
		return "0x4b7", nil // 1207 em hexadecimal
	}
	return "", errors.New("unsupported method")
}

// Call..
func (m MockedClientError) Call(method string, params []interface{}) (string, error) {
	return "", errors.New("RPC call failed")
}

// Call..
func (m MockedClientInvalidNumber) Call(method string, params []interface{}) (string, error) {
	return "invalid_number", nil
}
