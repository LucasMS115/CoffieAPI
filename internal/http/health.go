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
func (handler *HealthHandler) RegisterRoutes(serveMux *http.ServeMux) {
	serveMux.HandleFunc("GET /health", handler.Get)
}

// Get godoc
// @Summary Health check
// @Description Check if the API is running.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (handler *HealthHandler) Get(responseWriter http.ResponseWriter, request *http.Request) {
	_ = request
	response.JSON(responseWriter, http.StatusOK, map[string]string{"status": "ok"})
}
