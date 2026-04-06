package store

import (
	"context"
	"coffie/internal/feature/coffee/domain"
	"database/sql"
)

// PostgresCoffeeStore implements domain.CoffeeStore using PostgreSQL.
type PostgresCoffeeStore struct {
	db *sql.DB
}

// NewCoffeeStore creates a new Postgres-backed coffee store.
func NewCoffeeStore(db *sql.DB) *PostgresCoffeeStore {
	return &PostgresCoffeeStore{db: db}
}

func (s *PostgresCoffeeStore) Create(ctx context.Context, c *domain.Coffee) error {
	return nil
}

func (s *PostgresCoffeeStore) GetByID(ctx context.Context, id string) (*domain.Coffee, error) {
	return nil, nil
}

func (s *PostgresCoffeeStore) List(ctx context.Context, filter domain.ListFilter) ([]domain.Coffee, int, error) {
	return nil, 0, nil
}
