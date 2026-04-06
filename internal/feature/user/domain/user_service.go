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

// RegisterRequest holds validated input for creating a user.
type RegisterRequest struct {
	Name  string
	Email string
}

// UserStats holds computed user statistics.
type UserStats struct {
	RecipesCount   int
	AvgRatingGiven float64
	FavMethod      string
	FavCoffeeType  string
}

// UserStore defines the interface for user data persistence.
type UserStore interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetStats(ctx context.Context, userID string) (*UserStats, error)
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

// GetByID returns a user by ID.
func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	return nil, nil
}

// GetStats returns computed user stats.
func (s *Service) GetStats(ctx context.Context, id string) (*UserStats, error) {
	return nil, nil
}
