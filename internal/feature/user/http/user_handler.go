package http

import (
	"encoding/json"
	"net/http"

	"coffie/internal/feature/user/domain"
	"coffie/internal/http/response"
)

// Handler holds the HTTP handlers for user endpoints.
type Handler struct {
	userService *domain.Service
}

// NewHandler creates a new user handler.
func NewHandler(userService *domain.Service) *Handler {
	return &Handler{userService: userService}
}

// RegisterRoutes attaches user routes to the given ServeMux.
func (handler *Handler) RegisterRoutes(serveMux *http.ServeMux) {
	serveMux.HandleFunc("POST /api/users", handler.Register)
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account with name and email.
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterUser true "User registration"
// @Success 201 {object} UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users [post]
func (handler *Handler) Register(responseWriter http.ResponseWriter, request *http.Request) {
	// Parse request body
	var registerRequest RegisterUser
	if err := json.NewDecoder(request.Body).Decode(&registerRequest); err != nil {
		response.Error(responseWriter, http.StatusBadRequest, "INVALID_INPUT", "invalid request body", nil)
		return
	}

	// Validate
	validationErrors := ValidateRegisterUserRequest(&registerRequest)
	if len(validationErrors) > 0 {
		response.Error(responseWriter, http.StatusBadRequest, "INVALID_INPUT", "validation failed", validationErrors)
		return
	}

	// Adapt to domain
	serviceRequest := toRegisterRequest(&registerRequest)

	// Delegate to service
	createdUser, err := handler.userService.Register(request.Context(), serviceRequest)
	if err != nil {
		response.Error(responseWriter, http.StatusInternalServerError, "INTERNAL_ERROR", "could not create user", nil)
		return
	}

	// Build response
	userResponse := toUserResponse(createdUser)
	response.JSON(responseWriter, http.StatusCreated, userResponse)
}
