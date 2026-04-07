package http

import "coffie/internal/http/response"

func ValidateRegisterUserRequest(registerRequest *RegisterUser) []response.FieldError {
	validationErrors := []response.FieldError{}

	if registerRequest.Name == "" {
		validationErrors = append(validationErrors, response.FieldError{Field: "name", Message: "is required"})
	}

	if registerRequest.Email == "" {
		validationErrors = append(validationErrors, response.FieldError{Field: "email", Message: "is required"})
	}

	return validationErrors
}
