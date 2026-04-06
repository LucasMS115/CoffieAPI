package http

import (
	"net/http"

	"coffie/internal/feature/rating/domain"
)

// Handler holds the HTTP handlers for rating endpoints.
type Handler struct {
	ratingSvc *domain.Service
}

// NewHandler creates a new rating handler.
func NewHandler(ratingSvc *domain.Service) *Handler {
	return &Handler{ratingSvc: ratingSvc}
}

// RegisterRoutes attaches rating routes to the given ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// POST   /api/recipes/{id}/ratings
	// GET    /api/recipes/{id}/ratings
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) List(w http.ResponseWriter, r *http.Request)    {}
