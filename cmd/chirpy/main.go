package main

import (
	"log"
	"net/http"

	"github.com/kai-xlr/chirpy/internal/handlers"
	"github.com/kai-xlr/chirpy/internal/middleware"
)

func main() {
	const filepathRoot = "../.."
	const port = "8080"

	apiCfg := &handlers.APIConfig{}

	mux := http.NewServeMux()
	mux.Handle(
		"/app/",
		middleware.Metrics(apiCfg)(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))),
	)
	mux.HandleFunc("GET /api/healthz", handlers.Readiness)
	mux.HandleFunc("GET /api/metrics", middleware.MetricsHandler(apiCfg))
	mux.HandleFunc("POST /api/reset", apiCfg.Reset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
