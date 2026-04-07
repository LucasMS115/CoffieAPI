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
func (s *Service) Register(ctx context.Context, req RegisterRequest) (*User, error) {
	user := &User{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.userStore.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
