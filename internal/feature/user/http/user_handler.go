package http

import (
	"encoding/json"
	"net/http"

	"coffie/internal/feature/user/domain"
	"coffie/internal/http/response"
)

// Handler holds the HTTP handlers for user endpoints.
type Handler struct {
	userSvc *domain.Service
}

// NewHandler creates a new user handler.
func NewHandler(userSvc *domain.Service) *Handler {
	return &Handler{userSvc: userSvc}
}

// RegisterRoutes attaches user routes to the given ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", h.Register)
	// GET  /api/users/{id}
	// GET  /api/users/{id}/stats
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_INPUT", "invalid request body", nil)
		return
	}

	// Validate
	if req.Name == "" || req.Email == "" {
		fields := []response.FieldError{}
		if req.Name == "" {
			fields = append(fields, response.FieldError{Field: "name", Message: "is required"})
		}
		if req.Email == "" {
			fields = append(fields, response.FieldError{Field: "email", Message: "is required"})
		}
		response.Error(w, http.StatusBadRequest, "INVALID_INPUT", "validation failed", fields)
		return
	}

	// Adapt to domain
	svcReq := toRegisterRequest(&req)

	// Delegate to service
	created, err := h.userSvc.Register(r.Context(), svcReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", "could not create user", nil)
		return
	}

	// Build response
	resp := toUserResponse(created)
	response.JSON(w, http.StatusCreated, resp)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {}
