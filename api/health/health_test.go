package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a test HTTP request
	req := httptest.NewRequest("GET", "/health", nil)

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	// Validate the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res HealthResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate response data
	if res.Status != "OK" {
		t.Errorf("Expected status 'OK', got '%s'", res.Status)
	}
	if res.Message != "Service is running smoothly" {
		t.Errorf("Expected message 'Service is running smoothly', got '%s'", res.Message)
	}
}
