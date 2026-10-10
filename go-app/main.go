package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type AppResponse struct {
	Status      string    `json:"status"`
	Message     string    `json:"message"`
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
	Hostname    string    `json:"hostname"`
	Timestamp   time.Time `json:"timestamp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v2.0.0" // Updated lab version
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "docker-lab"
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-container"
	}

	// Root Endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request from %s on path %s", r.RemoteAddr, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		resp := AppResponse{
			Status:      "success",
			Message:     "Go application running smoothly inside Docker!",
			Version:     version,
			Environment: environment,
			Hostname:    hostname,
			Timestamp:   time.Now(),
		}
		json.NewEncoder(w).Encode(resp)
	})

	// Health Check Endpoint for Docker Probes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy","container":"%s"}`, hostname)
	})

	log.Printf("Starting Go Lab Service [%s] on port %s (Host: %s)...", version, port, hostname)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
