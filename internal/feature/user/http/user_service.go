package http

import (
	"context"

	"coffie/internal/feature/user/domain"
)

// UserService describes the user business operations required by the HTTP layer.
type UserService interface {
	Register(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error)
}