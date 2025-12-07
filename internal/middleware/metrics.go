package middleware

import (
	"fmt"
	"net/http"

	"github.com/kai-xlr/chirpy/internal/handlers"
)

func Metrics(cfg *handlers.APIConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cfg.FileserverHits.Add(1)
			next.ServeHTTP(w, r)
		})
	}
}

func MetricsHandler(cfg *handlers.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hits: %d", cfg.FileserverHits.Load())
	}
}
