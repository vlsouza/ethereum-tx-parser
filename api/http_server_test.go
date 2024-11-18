package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAddressFromRequest_ValidAddress(t *testing.T) {
	// Create a mock HTTP request with a valid address in the URL
	req := httptest.NewRequest("GET", "/some-endpoint/0x123", nil)
	// Use mux to set the path variables
	req = mux.SetURLVars(req, map[string]string{"address": "0x123"})

	address, err := getAddressFromRequest(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if address != "0x123" {
		t.Errorf("expected address '0x123', got '%s'", address)
	}
}

func TestGetAddressFromRequest_MissingAddress(t *testing.T) {
	// Create a mock HTTP request without an address in the URL
	req := httptest.NewRequest("GET", "/some-endpoint/", nil)
	// Use mux to set the path variables with an empty address
	req = mux.SetURLVars(req, map[string]string{"address": ""})

	address, err := getAddressFromRequest(req)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if address != "" {
		t.Errorf("expected empty address, got '%s'", address)
	}
}
