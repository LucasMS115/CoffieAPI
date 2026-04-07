package http

import (
	"net/http"

	"coffie/internal/feature/recipe/domain"
)

// Handler holds the HTTP handlers for recipe endpoints.
type Handler struct {
	recipeService *domain.Service
}

// NewHandler creates a new recipe handler.
func NewHandler(recipeService *domain.Service) *Handler {
	return &Handler{recipeService: recipeService}
}

// RegisterRoutes attaches recipe routes to the given ServeMux.
func (handler *Handler) RegisterRoutes(serveMux *http.ServeMux) {
	// POST   /api/recipes
	// GET    /api/recipes/{id}
	// GET    /api/recipes
	// PUT    /api/recipes/{id}
	// DELETE /api/recipes/{id}
	_ = handler
	_ = serveMux
}

func (handler *Handler) Create(responseWriter http.ResponseWriter, request *http.Request) {}
func (handler *Handler) Get(responseWriter http.ResponseWriter, request *http.Request)    {}
func (handler *Handler) List(responseWriter http.ResponseWriter, request *http.Request)   {}
func (handler *Handler) Update(responseWriter http.ResponseWriter, request *http.Request) {}
func (handler *Handler) Delete(responseWriter http.ResponseWriter, request *http.Request) {}
