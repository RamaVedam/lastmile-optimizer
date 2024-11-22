package routing

import (
	"encoding/json"
	"net/http"
)

type PolygonRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance_km"`
}

type PolygonResponse struct {
	Polygon string `json:"polygon"`
}

func GeneratePolygonHandler(w http.ResponseWriter, r *http.Request) {
	var req PolygonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Mocked polygon generation (replace with actual logic later)
	mockPolygon := `{"type":"Polygon","coordinates":[[[13.405,52.52],[13.41,52.52],[13.41,52.525],[13.405,52.525],[13.405,52.52]]]}`
	res := PolygonResponse{Polygon: mockPolygon}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
