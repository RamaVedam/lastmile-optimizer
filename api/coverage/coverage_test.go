package coverage

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratePolygonHandler(t *testing.T) {
	// Mock request payload
	payload := PolygonRequest{
		HubLatitude:   52.5200,
		HubLongitude:  13.4050,
		PurchaseValue: 50,
		TimeMinutes:   10,
	}
	reqBody, _ := json.Marshal(payload)

	// Create a test HTTP request
	req := httptest.NewRequest("POST", "/coverage", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GeneratePolygonHandler)
	handler.ServeHTTP(rr, req)

	// Validate the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res PolygonResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate response data
	if res.VehicleType != "scooter" {
		t.Errorf("Expected vehicle type 'scooter', got '%s'", res.VehicleType)
	}
	if res.Polygon == "" {
		t.Error("Expected polygon data, got empty string")
	}
}
