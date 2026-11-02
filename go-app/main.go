package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Service     string `json:"service"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	response := HealthResponse{
		Status:      "healthy",
		Environment: env,
		Service:     "go-app",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go application starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
