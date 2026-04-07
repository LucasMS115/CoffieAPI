package store

import (
	"context"
	"database/sql"

	"coffie/internal/feature/coffee/domain"
)

// PostgresCoffeeStore implements domain.CoffeeStore using PostgreSQL.
type PostgresCoffeeStore struct {
	databaseConnection *sql.DB
}

// NewCoffeeStore creates a new Postgres-backed coffee store.
func NewCoffeeStore(databaseConnection *sql.DB) *PostgresCoffeeStore {
	return &PostgresCoffeeStore{databaseConnection: databaseConnection}
}

func (postgresCoffeeStore *PostgresCoffeeStore) Create(requestContext context.Context, coffee *domain.Coffee) error {
	_ = postgresCoffeeStore
	_ = requestContext
	_ = coffee
	return nil
}

func (postgresCoffeeStore *PostgresCoffeeStore) GetByID(requestContext context.Context, coffeeID string) (*domain.Coffee, error) {
	_ = postgresCoffeeStore
	_ = requestContext
	_ = coffeeID
	return nil, nil
}

func (postgresCoffeeStore *PostgresCoffeeStore) List(requestContext context.Context, filter domain.ListFilter) ([]domain.Coffee, int, error) {
	_ = postgresCoffeeStore
	_ = requestContext
	_ = filter
	return nil, 0, nil
}
