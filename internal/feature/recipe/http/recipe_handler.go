package http

import (
	"net/http"

	"coffie/internal/feature/recipe/domain"
)

// Handler holds the HTTP handlers for recipe endpoints.
type Handler struct {
	recipeSvc *domain.Service
}

// NewHandler creates a new recipe handler.
func NewHandler(recipeSvc *domain.Service) *Handler {
	return &Handler{recipeSvc: recipeSvc}
}

// RegisterRoutes attaches recipe routes to the given ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// POST   /api/recipes
	// GET    /api/recipes/{id}
	// GET    /api/recipes
	// PUT    /api/recipes/{id}
	// DELETE /api/recipes/{id}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) List(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request)  {}
