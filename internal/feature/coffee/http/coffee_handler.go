package http

import (
	"net/http"

	"coffie/internal/feature/coffee/domain"
)

// Handler holds the HTTP handlers for coffee endpoints.
type Handler struct {
	coffeeSvc *domain.Service
}

// NewHandler creates a new coffee handler.
func NewHandler(coffeeSvc *domain.Service) *Handler {
	return &Handler{coffeeSvc: coffeeSvc}
}

// RegisterRoutes attaches coffee routes to the given ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// POST /api/coffees
	// GET  /api/coffees/{id}
	// GET  /api/coffees
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) List(w http.ResponseWriter, r *http.Request)    {}
