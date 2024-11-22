package main

import (
	"log"
	"net/http"

	"lastmile-optimizer/api/coverage"
	"lastmile-optimizer/api/health"
	"lastmile-optimizer/api/route"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register API endpoints
	r.HandleFunc("/coverage", coverage.GeneratePolygonHandler).Methods("POST")
	r.HandleFunc("/route", route.OptimizeRouteHandler).Methods("POST")
	r.HandleFunc("/health", health.HealthCheckHandler).Methods("GET")

	// Start the HTTP server
	port := "8080"
	log.Printf("Starting Route Optimizer service on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
