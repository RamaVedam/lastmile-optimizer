package health

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the response format for health checks
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthCheckHandler handles the health check endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "OK",
		Message: "Service is running smoothly",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
