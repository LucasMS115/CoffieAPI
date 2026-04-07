package domain

import (
	"context"
	"errors"
)

// RecipeStore defines the interface for recipe data persistence.
type RecipeStore interface {
	Create(requestContext context.Context, recipe *Recipe) error
	GetByID(requestContext context.Context, recipeID string) (*RecipeWithDetails, error)
	List(requestContext context.Context, filter ListFilter) ([]RecipeSummary, int, error)
	Update(requestContext context.Context, recipe *Recipe) error
	Delete(requestContext context.Context, recipeID string) error
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
func (service *Service) Create(requestContext context.Context, createRecipeRequest CreateRecipeRequest) (*Recipe, error) {
	_ = requestContext
	_ = createRecipeRequest
	return nil, nil
}

// GetByID returns a recipe with full details.
func (service *Service) GetByID(requestContext context.Context, recipeID string) (*RecipeWithDetails, error) {
	_ = service
	_ = requestContext
	_ = recipeID
	return nil, nil
}

// List searches and lists recipes.
func (service *Service) List(requestContext context.Context, listRecipesRequest ListRecipesRequest) ([]RecipeSummary, int, error) {
	_ = service
	_ = requestContext
	_ = listRecipesRequest
	return nil, 0, nil
}

// Update updates an existing recipe.
func (service *Service) Update(requestContext context.Context, recipeID string, updateRecipeRequest UpdateRecipeRequest) (*Recipe, error) {
	_ = service
	_ = requestContext
	_ = recipeID
	_ = updateRecipeRequest
	return nil, nil
}

// Delete removes a recipe.
func (service *Service) Delete(requestContext context.Context, recipeID string) error {
	_ = service
	_ = requestContext
	_ = recipeID
	return nil
}
