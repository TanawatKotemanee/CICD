package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"simple-go-api/config"
)

func main() {
	// Load environment variables`
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize the database
	config.InitDB()
	if config.DB != nil {
		defer config.DB.Close()
	}

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", getHealthCheckHandler)
	mux.HandleFunc("GET /db", getDBVersionHandler)
	mux.HandleFunc("GET /version", getVersionHandler)

	port := os.Getenv("PORT")
	log.Printf("Server started on port %s\n", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func getHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status": "Healthy",
	}); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func getDBVersionHandler(w http.ResponseWriter, r *http.Request) {
	if config.DB == nil {
		http.Error(w, "Database connection is not established", http.StatusInternalServerError)
		return
	}
	var version string
	err := config.DB.Get(&version, "SELECT VERSION()")
	if err != nil {
		http.Error(w, "Failed to get database version", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"version": version,
	}); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func getVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"version": "1.0.1",
	}); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
