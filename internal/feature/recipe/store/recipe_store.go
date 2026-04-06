package store

import (
	"context"
	"coffie/internal/feature/recipe/domain"
	"database/sql"
)

// PostgresRecipeStore implements domain.RecipeStore using PostgreSQL.
type PostgresRecipeStore struct {
	db *sql.DB
}

// NewRecipeStore creates a new Postgres-backed recipe store.
func NewRecipeStore(db *sql.DB) *PostgresRecipeStore {
	return &PostgresRecipeStore{db: db}
}

func (s *PostgresRecipeStore) Create(ctx context.Context, r *domain.Recipe) error {
	return nil
}

func (s *PostgresRecipeStore) GetByID(ctx context.Context, id string) (*domain.RecipeWithDetails, error) {
	return nil, nil
}

func (s *PostgresRecipeStore) List(ctx context.Context, filter domain.ListFilter) ([]domain.RecipeSummary, int, error) {
	return nil, 0, nil
}

func (s *PostgresRecipeStore) Update(ctx context.Context, r *domain.Recipe) error {
	return nil
}

func (s *PostgresRecipeStore) Delete(ctx context.Context, id string) error {
	return nil
}
