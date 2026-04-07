package http

import "coffie/internal/http/response"

// RegisterUser holds the validated body for POST /api/users.
type RegisterUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate checks whether the user registration payload is complete.
func (registerRequest *RegisterUser) Validate() []response.FieldError {
	validationErrors := []response.FieldError{}

	if registerRequest.Name == "" {
		validationErrors = append(validationErrors, response.FieldError{Field: "name", Message: "is required"})
	}

	if registerRequest.Email == "" {
		validationErrors = append(validationErrors, response.FieldError{Field: "email", Message: "is required"})
	}

	return validationErrors
}
