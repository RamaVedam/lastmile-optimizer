package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs details about each HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r) // Call the next middleware/handler in the chain

		duration := time.Since(start)
		log.Printf("Completed %s in %v", r.URL.Path, duration)
	})
}
