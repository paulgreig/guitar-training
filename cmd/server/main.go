package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulgreig/guitar-training/internal/api"
	"github.com/paulgreig/guitar-training/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize router
	router := api.SetupRouter()

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	log.Printf("API endpoints available at http://localhost%s/api", serverAddr)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
