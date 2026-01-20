package impact_partner

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Controller handles HTTP requests for impact partners
type Controller struct {
	service Service
}

// NewController creates a new controller
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

// HandleGetAll handles GET /api/impact-partners
func (c *Controller) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	partners, err := c.service.GetAllPartners(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partners)
}

// HandleGetByID handles GET /api/impact-partners/{id}
func (c *Controller) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/impact-partners/")
	id := strings.Split(path, "/")[0]

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	partner, err := c.service.GetPartnerByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partner)
}
