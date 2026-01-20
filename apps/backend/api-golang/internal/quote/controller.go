package quote

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Controller handles HTTP requests for quotes
type Controller struct {
	orchestrator *Orchestrator
}

// NewController creates a new quote controller
func NewController(orchestrator *Orchestrator) *Controller {
	return &Controller{orchestrator: orchestrator}
}

// HandleCreateQuote handles POST /api/quotes
func (c *Controller) HandleCreateQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		c.writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	// Parse request body
	var req CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body: "+err.Error())
		return
	}

	// Debug: Check if there's an unexpected transactionId field
	// This shouldn't be in the request, but if it is, we'll ignore it

	// Get organisation ID from header (simulating auth)
	headerOrgID := r.Header.Get("X-Organisation-ID")
	if headerOrgID == "" {
		// Default to the organisation ID in the body for demo purposes
		headerOrgID = req.OrganisationID
	}

	// Set defaults
	if req.Locale == "" {
		req.Locale = "en-GB"
	}

	// Validate required fields
	if req.OrganisationID == "" {
		c.writeError(w, http.StatusBadRequest, "MISSING_FIELD", "organisationId is required")
		return
	}
	if req.Customer.Reference == "" {
		c.writeError(w, http.StatusBadRequest, "MISSING_FIELD", "customer.reference is required")
		return
	}
	if req.Customer.Country == "" {
		c.writeError(w, http.StatusBadRequest, "MISSING_FIELD", "customer.country is required")
		return
	}

	// Note: transactionId is NOT required - it's auto-generated in the orchestrator
	// Validate that we have either orderItems or some way to determine transaction amount
	if len(req.OrderItems) == 0 {
		// For demo purposes, we'll allow empty orderItems and use a default
		// In production, this might be an error or require a transaction amount field
	}

	// Create quote
	ctx := r.Context()
	response, err := c.orchestrator.CreateQuote(ctx, &req, headerOrgID)
	if err != nil {
		// Check for specific error types
		errStr := err.Error()
		switch {
		case strings.Contains(errStr, "forbidden") || strings.Contains(errStr, "FORBIDDEN"):
			c.writeError(w, http.StatusForbidden, "FORBIDDEN", err.Error())
		case strings.Contains(errStr, "not found"):
			c.writeError(w, http.StatusNotFound, "NOT_FOUND", err.Error())
		case strings.Contains(errStr, "validation"):
			c.writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		default:
			c.writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		}
		return
	}

	c.writeJSON(w, http.StatusCreated, response)
}

// HandleGetQuote handles GET /api/quotes/{id}
func (c *Controller) HandleGetQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		c.writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	// Extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/quotes/")
	id := strings.Split(path, "/")[0]

	if id == "" {
		c.writeError(w, http.StatusBadRequest, "MISSING_FIELD", "Quote ID is required")
		return
	}

	ctx := r.Context()
	quote, err := c.orchestrator.GetQuote(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.writeError(w, http.StatusNotFound, "NOT_FOUND", "Quote not found")
		} else {
			c.writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		}
		return
	}

	c.writeJSON(w, http.StatusOK, quote)
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *Controller) writeError(w http.ResponseWriter, status int, code, message string) {
	resp := ErrorResponse{}
	resp.Error.Code = code
	resp.Error.Message = message
	c.writeJSON(w, status, resp)
}

func (c *Controller) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
