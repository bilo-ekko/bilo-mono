package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func main() {
	// Register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)

	// Start server
	port := ":9000"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  - GET http://localhost" + port + "/")
	fmt.Println("  - GET http://localhost" + port + "/api/health")
	fmt.Println("  - GET http://localhost" + port + "/api/hello")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// homeHandler handles requests to the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to api-golang!",
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

// healthHandler handles health check requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

// helloHandler handles hello requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}
