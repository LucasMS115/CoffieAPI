package http

import (
	"net/http"

	"coffie/internal/feature/rating/domain"
)

// Handler holds the HTTP handlers for rating endpoints.
type Handler struct {
	ratingService *domain.Service
}

// NewHandler creates a new rating handler.
func NewHandler(ratingService *domain.Service) *Handler {
	return &Handler{ratingService: ratingService}
}

// RegisterRoutes attaches rating routes to the given ServeMux.
func (handler *Handler) RegisterRoutes(serveMux *http.ServeMux) {
	// POST   /api/recipes/{id}/ratings
	// GET    /api/recipes/{id}/ratings
	_ = handler
	_ = serveMux
}

func (handler *Handler) Create(responseWriter http.ResponseWriter, request *http.Request) {}
func (handler *Handler) List(responseWriter http.ResponseWriter, request *http.Request)   {}
