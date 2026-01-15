package main

import (
	"api-golang/internal/impact_partner"
	"api-golang/internal/impact_project"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bilo-mono/packages/common/calculator"
	"github.com/bilo-mono/packages/common/logger"
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
	// Initialize shared logger
	appLogger := logger.NewLogger("Go-API")
	appLogger.Info("Initializing Go API application...")

	// Demonstrate calculator usage
	calcResult := calculator.CalculateX(calculator.CalculateXInput{
		Value:      100.0,
		Multiplier: ptrFloat64(1.5),
		Offset:     ptrFloat64(25.0),
	})
	appLogger.Infof("Calculation result: %s", calcResult.Formula)

	// Initialize Impact Partner module
	partnerRepo := impact_partner.NewRepository()
	partnerService := impact_partner.NewService(partnerRepo)
	partnerController := impact_partner.NewController(partnerService)

	// Initialize Impact Project module
	projectRepo := impact_project.NewRepository()
	projectService := impact_project.NewService(projectRepo)
	projectController := impact_project.NewController(projectService)

	// Register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)

	// Impact Partners routes
	http.HandleFunc("/api/impact-partners", partnerController.HandleGetAll)
	http.HandleFunc("/api/impact-partners/", func(w http.ResponseWriter, r *http.Request) {
		// Route to specific partner if ID is present
		if strings.TrimPrefix(r.URL.Path, "/api/impact-partners/") != "" {
			partnerController.HandleGetByID(w, r)
		} else {
			partnerController.HandleGetAll(w, r)
		}
	})

	// Impact Projects routes
	http.HandleFunc("/api/impact-projects", projectController.HandleGetAll)
	http.HandleFunc("/api/impact-projects/", func(w http.ResponseWriter, r *http.Request) {
		// Route to specific project if ID is present
		if strings.TrimPrefix(r.URL.Path, "/api/impact-projects/") != "" {
			projectController.HandleGetByID(w, r)
		} else {
			projectController.HandleGetAll(w, r)
		}
	})

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Go API Server starting on http://localhost%s\n", port)
	fmt.Println("\n📍 Available endpoints:")
	fmt.Println("  - GET  http://localhost" + port + "/")
	fmt.Println("  - GET  http://localhost" + port + "/api/health")
	fmt.Println("  - GET  http://localhost" + port + "/api/hello?name=World")
	fmt.Println("\n🌍 Impact Partners:")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-partners")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-partners/{id}")
	fmt.Println("\n🌱 Impact Projects:")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects/{id}")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects?partnerId={id}")
	fmt.Println()

	appLogger.Infof("Server listening on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		appLogger.Error("Server failed to start", err)
	}
}

// ptrFloat64 returns a pointer to a float64 value
func ptrFloat64(v float64) *float64 {
	return &v
}

// homeHandler handles requests to the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to api-golang! 🚀",
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
