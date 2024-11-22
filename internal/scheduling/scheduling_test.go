package scheduling

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestScheduleTaskHandler(t *testing.T) {
	reqBody := `{"task_id": 1, "delivery_time": "2023-12-01T10:00:00Z", "product_type": "cold", "weight_kg": 5, "vehicle_type": "car"}`
	req := httptest.NewRequest("POST", "/schedule", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ScheduleTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"SUCCESS","message":"Task 1 scheduled for delivery at 2023-12-01T10:00:00Z using vehicle car"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
