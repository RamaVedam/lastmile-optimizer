package routing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratePolygonHandler(t *testing.T) {
	reqBody := `{"latitude":52.52,"longitude":13.405,"distance_km":5}`
	req := httptest.NewRequest("POST", "/coverage", bytes.NewReader([]byte(reqBody)))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GeneratePolygonHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res PolygonResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if res.Polygon == "" {
		t.Error("Expected polygon data, got empty string")
	}
}
