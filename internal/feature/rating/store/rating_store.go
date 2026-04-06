package store

import (
	"context"
	"coffie/internal/feature/rating/domain"
	"database/sql"
)

// PostgresRatingStore implements domain.RatingStore using PostgreSQL.
type PostgresRatingStore struct {
	db *sql.DB
}

// NewRatingStore creates a new Postgres-backed rating store.
func NewRatingStore(db *sql.DB) *PostgresRatingStore {
	return &PostgresRatingStore{db: db}
}

func (s *PostgresRatingStore) Create(ctx context.Context, r *domain.Rating) error {
	return nil
}

func (s *PostgresRatingStore) ListByRecipeID(ctx context.Context, recipeID string) ([]domain.Rating, error) {
	return nil, nil
}

func (s *PostgresRatingStore) GetAvgByRecipeID(ctx context.Context, recipeID string) (*float64, int, error) {
	return nil, 0, nil
}
