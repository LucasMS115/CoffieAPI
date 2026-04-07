package domain

import (
	"context"
	"errors"
)

// RatingStore defines the interface for rating data persistence.
type RatingStore interface {
	Create(requestContext context.Context, rating *Rating) error
	ListByRecipeID(requestContext context.Context, recipeID string) ([]Rating, error)
	GetAvgByRecipeID(requestContext context.Context, recipeID string) (*float64, int, error)
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
func (service *Service) Create(requestContext context.Context, createRatingRequest CreateRatingRequest) (*Rating, error) {
	_ = requestContext
	_ = createRatingRequest
	return nil, nil
}

// ListByRecipeID returns all ratings for a recipe.
func (service *Service) ListByRecipeID(requestContext context.Context, recipeID string) ([]Rating, error) {
	_ = service
	_ = requestContext
	_ = recipeID
	return nil, nil
}
