package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOptimizeRouteHandler(t *testing.T) {
	// Mock request payload
	payload := RouteRequest{
		StartVertex: 1,
		EndVertex:   10,
	}
	reqBody, _ := json.Marshal(payload)

	// Create a test HTTP request
	req := httptest.NewRequest("POST", "/route", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(OptimizeRouteHandler)
	handler.ServeHTTP(rr, req)

	// Validate the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res RouteResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate response data
	if res.Route == "" {
		t.Error("Expected route data, got empty string")
	}
	if res.TotalTime != 12.5 {
		t.Errorf("Expected total time 12.5, got %f", res.TotalTime)
	}
}
