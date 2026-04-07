package response

import (
	"errors"
	"net/http"

	userdomain "coffie/internal/feature/user/domain"
)

// DomainError writes a standardized HTTP response for known domain errors.
func DomainError(responseWriter http.ResponseWriter, domainError error) {
	switch {
	case errors.Is(domainError, userdomain.ErrUserAlreadyExists):
		Error(responseWriter, http.StatusConflict, "CONFLICT", "user already exists", nil)
	default:
		Error(responseWriter, http.StatusInternalServerError, "INTERNAL_ERROR", "internal server error", nil)
	}
}