package store

import (
	"context"
	"database/sql"

	"coffie/internal/feature/rating/domain"
)

// PostgresRatingStore implements domain.RatingStore using PostgreSQL.
type PostgresRatingStore struct {
	databaseConnection *sql.DB
}

// NewRatingStore creates a new Postgres-backed rating store.
func NewRatingStore(databaseConnection *sql.DB) *PostgresRatingStore {
	return &PostgresRatingStore{databaseConnection: databaseConnection}
}

func (postgresRatingStore *PostgresRatingStore) Create(requestContext context.Context, rating *domain.Rating) error {
	_ = postgresRatingStore
	_ = requestContext
	_ = rating
	return nil
}

func (postgresRatingStore *PostgresRatingStore) ListByRecipeID(requestContext context.Context, recipeID string) ([]domain.Rating, error) {
	_ = postgresRatingStore
	_ = requestContext
	_ = recipeID
	return nil, nil
}

func (postgresRatingStore *PostgresRatingStore) GetAvgByRecipeID(requestContext context.Context, recipeID string) (*float64, int, error) {
	_ = postgresRatingStore
	_ = requestContext
	_ = recipeID
	return nil, 0, nil
}
