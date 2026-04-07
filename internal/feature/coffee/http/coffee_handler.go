package http

import (
	"net/http"

	"coffie/internal/feature/coffee/domain"
)

// Handler holds the HTTP handlers for coffee endpoints.
type Handler struct {
	coffeeService *domain.Service
}

// NewHandler creates a new coffee handler.
func NewHandler(coffeeService *domain.Service) *Handler {
	return &Handler{coffeeService: coffeeService}
}

// RegisterRoutes attaches coffee routes to the given ServeMux.
func (handler *Handler) RegisterRoutes(serveMux *http.ServeMux) {
	// POST /api/coffees
	// GET  /api/coffees/{id}
	// GET  /api/coffees
	_ = handler
	_ = serveMux
}

func (handler *Handler) Create(responseWriter http.ResponseWriter, request *http.Request) {}
func (handler *Handler) Get(responseWriter http.ResponseWriter, request *http.Request)    {}
func (handler *Handler) List(responseWriter http.ResponseWriter, request *http.Request)   {}
