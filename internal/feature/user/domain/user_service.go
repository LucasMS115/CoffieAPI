package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Service handles user business logic.
type Service struct {
	userStore UserStore
}

// NewService creates a new Service.
func NewService(userStore UserStore) *Service {
	return &Service{userStore: userStore}
}

// Register handles user registration.
func (service *Service) Register(requestContext context.Context, registerRequest RegisterRequest) (*User, error) {
	user := &User{
		ID:        uuid.NewString(),
		Name:      registerRequest.Name,
		Email:     registerRequest.Email,
		CreatedAt: time.Now().UTC(),
	}

	if createError := service.userStore.Create(requestContext, user); createError != nil {
		return nil, createError
	}

	return user, nil
}
