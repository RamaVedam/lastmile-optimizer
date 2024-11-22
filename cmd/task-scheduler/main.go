package main

import (
	"log"
	"net/http"

	"lastmile-optimizer/internal/monitoring"
	"lastmile-optimizer/internal/scheduling"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task Scheduler Service is healthy"))
	}).Methods("GET")

	// Endpoint for scheduling tasks
	r.HandleFunc("/schedule", scheduling.ScheduleTaskHandler).Methods("POST")

	// Prometheus metrics endpoint
	monitoring.InitMonitoring()
	r.Handle("/metrics", http.HandlerFunc(monitoring.MetricsHandler))

	// Start the HTTP server
	log.Println("Starting Task Scheduler service on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
