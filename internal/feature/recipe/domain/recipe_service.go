package domain

import (
	"context"
	"errors"
)

// RecipeStore defines the interface for recipe data persistence.
type RecipeStore interface {
	Create(ctx context.Context, r *Recipe) error
	GetByID(ctx context.Context, id string) (*RecipeWithDetails, error)
	List(ctx context.Context, filter ListFilter) ([]RecipeSummary, int, error)
	Update(ctx context.Context, r *Recipe) error
	Delete(ctx context.Context, id string) error
}

// ListFilter holds optional filters for listing recipes.
type ListFilter struct {
	UserID string
	Method string
	Search string
	Page   int
	Limit  int
}

// Service handles recipe business logic.
type Service struct {
	recipeStore RecipeStore
}

// NewService creates a new Service.
func NewService(recipeStore RecipeStore) *Service {
	return &Service{recipeStore: recipeStore}
}

var (
	ErrRecipeNotFound = errors.New("recipe not found")
	ErrUnauthorized   = errors.New("not authorized to modify this recipe")
)

// CreateRecipeRequest holds validated input for creating a recipe.
type CreateRecipeRequest struct {
	UserID      string
	CoffeeID    string
	Method      string
	WaterTemp   int
	Dose        float64
	Yield       float64
	BrewTime    int
	Description string
}

// UpdateRecipeRequest holds validated input for updating a recipe.
type UpdateRecipeRequest struct {
	Method      *string
	WaterTemp   *int
	Dose        *float64
	Yield       *float64
	BrewTime    *int
	Description *string
}

// ListRecipesRequest holds validated query params for listing recipes.
type ListRecipesRequest struct {
	UserID string
	Method string
	Search string
	Page   int
	Limit  int
}

// Create creates a new recipe.
func (s *Service) Create(ctx context.Context, req CreateRecipeRequest) (*Recipe, error) {
	return nil, nil
}

// GetByID returns a recipe with full details.
func (s *Service) GetByID(ctx context.Context, id string) (*RecipeWithDetails, error) {
	return nil, nil
}

// List searches and lists recipes.
func (s *Service) List(ctx context.Context, req ListRecipesRequest) ([]RecipeSummary, int, error) {
	return nil, 0, nil
}

// Update updates an existing recipe.
func (s *Service) Update(ctx context.Context, id string, req UpdateRecipeRequest) (*Recipe, error) {
	return nil, nil
}

// Delete removes a recipe.
func (s *Service) Delete(ctx context.Context, id string) error {
	return nil
}
