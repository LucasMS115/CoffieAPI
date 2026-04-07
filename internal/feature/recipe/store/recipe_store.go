package store

import (
	"context"
	"database/sql"

	"coffie/internal/feature/recipe/domain"
)

// PostgresRecipeStore implements domain.RecipeStore using PostgreSQL.
type PostgresRecipeStore struct {
	databaseConnection *sql.DB
}

// NewRecipeStore creates a new Postgres-backed recipe store.
func NewRecipeStore(databaseConnection *sql.DB) *PostgresRecipeStore {
	return &PostgresRecipeStore{databaseConnection: databaseConnection}
}

func (postgresRecipeStore *PostgresRecipeStore) Create(requestContext context.Context, recipe *domain.Recipe) error {
	_ = postgresRecipeStore
	_ = requestContext
	_ = recipe
	return nil
}

func (postgresRecipeStore *PostgresRecipeStore) GetByID(requestContext context.Context, recipeID string) (*domain.RecipeWithDetails, error) {
	_ = postgresRecipeStore
	_ = requestContext
	_ = recipeID
	return nil, nil
}

func (postgresRecipeStore *PostgresRecipeStore) List(requestContext context.Context, filter domain.ListFilter) ([]domain.RecipeSummary, int, error) {
	_ = postgresRecipeStore
	_ = requestContext
	_ = filter
	return nil, 0, nil
}

func (postgresRecipeStore *PostgresRecipeStore) Update(requestContext context.Context, recipe *domain.Recipe) error {
	_ = postgresRecipeStore
	_ = requestContext
	_ = recipe
	return nil
}

func (postgresRecipeStore *PostgresRecipeStore) Delete(requestContext context.Context, recipeID string) error {
	_ = postgresRecipeStore
	_ = requestContext
	_ = recipeID
	return nil
}
