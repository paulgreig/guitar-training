package api

import (
	"net/http"

	"github.com/paulgreig/guitar-training/internal/api/handlers"
)

// SetupRouter configures and returns the HTTP router
func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	// API routes
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/health", handlers.HealthCheck)
	apiMux.HandleFunc("/chords", handlers.ChordHandler)
	apiMux.HandleFunc("/scales", handlers.ScaleHandler)

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return mux
}
