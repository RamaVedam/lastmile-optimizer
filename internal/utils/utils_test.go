package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	mockData := map[string]string{"message": "success"}

	WriteJSONResponse(rr, http.StatusOK, mockData)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res map[string]string
	err := json.NewDecoder(rr.Body).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if res["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", res["message"])
	}
}
