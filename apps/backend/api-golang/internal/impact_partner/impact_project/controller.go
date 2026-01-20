package impact_project

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Controller handles HTTP requests for impact projects
type Controller struct {
	service *Service
}

// NewController creates a new controller
func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

// HandleGetAll handles GET /api/impact-projects
func (c *Controller) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if filtering by partner
	partnerID := r.URL.Query().Get("partnerId")
	if partnerID != "" {
		projects := c.service.GetProjectsByPartnerID(partnerID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
		return
	}

	projects := c.service.GetAllProjects()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// HandleGetByID handles GET /api/impact-projects/{id}
func (c *Controller) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/impact-projects/")
	id := strings.Split(path, "/")[0]

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	project, err := c.service.GetProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}
