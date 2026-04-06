package domain

import (
	"context"
	"errors"
)

// RatingStore defines the interface for rating data persistence.
type RatingStore interface {
	Create(ctx context.Context, r *Rating) error
	ListByRecipeID(ctx context.Context, recipeID string) ([]Rating, error)
	GetAvgByRecipeID(ctx context.Context, recipeID string) (*float64, int, error)
}

// Service handles rating business logic.
type Service struct {
	ratingStore RatingStore
}

// NewService creates a new Service.
func NewService(ratingStore RatingStore) *Service {
	return &Service{ratingStore: ratingStore}
}

var (
	ErrRecipeAlreadyRated = errors.New("user already rated this recipe")
	ErrInvalidScore       = errors.New("score must be between 1 and 5")
)

// CreateRatingRequest holds validated input for creating a rating.
type CreateRatingRequest struct {
	RecipeID string
	UserID   string
	Score    int
	Comment  string
}

// Create submits a rating for a recipe.
func (s *Service) Create(ctx context.Context, req CreateRatingRequest) (*Rating, error) {
	return nil, nil
}

// ListByRecipeID returns all ratings for a recipe.
func (s *Service) ListByRecipeID(ctx context.Context, recipeID string) ([]Rating, error) {
	return nil, nil
}
