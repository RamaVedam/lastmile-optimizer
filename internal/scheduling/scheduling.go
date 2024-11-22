package scheduling

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TaskRequest represents the input payload for scheduling a task
type TaskRequest struct {
	TaskID       int     `json:"task_id"`
	DeliveryTime string  `json:"delivery_time"`
	ProductType  string  `json:"product_type"`
	WeightKg     float64 `json:"weight_kg"`
	VehicleType  string  `json:"vehicle_type"`
}

// TaskResponse represents the task scheduling output
type TaskResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ScheduleTaskHandler handles task scheduling requests
func ScheduleTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON payload
	var req TaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Placeholder logic for task scheduling
	message := fmt.Sprintf("Task %d scheduled for delivery at %s using vehicle %s", req.TaskID, req.DeliveryTime, req.VehicleType)
	status := "SUCCESS"

	// Respond with the task scheduling status
	res := TaskResponse{
		Status:  status,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
