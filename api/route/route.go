package route

import (
	"encoding/json"
	"net/http"
)

type RouteRequest struct {
	StartVertex int `json:"start_vertex"`
	EndVertex   int `json:"end_vertex"`
}

type RouteResponse struct {
	Route     string  `json:"route"`
	TotalTime float64 `json:"total_time"`
}

func OptimizeRouteHandler(w http.ResponseWriter, r *http.Request) {
	var req RouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mockRoute := `{"type":"LineString","coordinates":[[13.405,52.52],[13.41,52.525]]}`
	mockTime := 12.5

	res := RouteResponse{
		Route:     mockRoute,
		TotalTime: mockTime,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
