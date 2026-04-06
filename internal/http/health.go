package http

import (
	"net/http"

	"coffie/internal/http/response"
)

// HealthHandler handles the /health endpoint.
type HealthHandler struct{}

// NewHealthHandler creates a new health check handler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// RegisterRoutes attaches the health route to the given ServeMux.
func (h *HealthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.Get)
}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
