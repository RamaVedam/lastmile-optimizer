package auth

import (
	"net/http"
	"strings"
)

// APIKey represents a placeholder for API key validation (can be dynamic or fetched from env/DB)
const APIKey = "my-secure-api-key"

// AuthenticateAPIKey is middleware to validate API key in request headers
func AuthenticateAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		if !strings.HasPrefix(apiKey, "Bearer ") || strings.TrimPrefix(apiKey, "Bearer ") != APIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
