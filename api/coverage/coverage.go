package coverage

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PolygonRequest struct {
	HubLatitude   float64 `json:"hub_latitude"`
	HubLongitude  float64 `json:"hub_longitude"`
	PurchaseValue float64 `json:"purchase_value"`
	TimeMinutes   float64 `json:"time_minutes"`
}

type PolygonResponse struct {
	VehicleType string `json:"vehicle_type"`
	Polygon     string `json:"polygon"`
}

func GeneratePolygonHandler(w http.ResponseWriter, r *http.Request) {
	var req PolygonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vehicleType, polygon, err := calculatePolygon(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := PolygonResponse{
		VehicleType: vehicleType,
		Polygon:     polygon,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func calculatePolygon(req PolygonRequest) (string, string, error) {
	vehicleType, err := determineVehicle(req.PurchaseValue)
	if err != nil {
		return "", "", err
	}

	mockPolygon := `{"type":"Polygon","coordinates":[[[13.405,52.52],[13.41,52.52],[13.41,52.525],[13.405,52.525],[13.405,52.52]]]}`
	return vehicleType, mockPolygon, nil
}

func determineVehicle(purchaseValue float64) (string, error) {
	switch {
	case purchaseValue <= 25:
		return "bicycle", nil
	case purchaseValue <= 100:
		return "scooter", nil
	case purchaseValue <= 200:
		return "car", nil
	default:
		return "", errors.New("purchase value exceeds delivery limits")
	}
}
