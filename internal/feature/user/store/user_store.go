package store

import (
	"context"
	"database/sql"

	"coffie/internal/feature/user/domain"
)

// PostgresUserStore implements domain.UserStore using PostgreSQL.
type PostgresUserStore struct {
	databaseConnection *sql.DB
}

// NewUserStore creates a new Postgres-backed user store.
func NewUserStore(databaseConnection *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{databaseConnection: databaseConnection}
}

func (postgresUserStore *PostgresUserStore) Create(requestContext context.Context, user *domain.User) error {
	const query = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	if _, executeError := postgresUserStore.databaseConnection.ExecContext(requestContext, query, user.ID, user.Name, user.Email, user.CreatedAt); executeError != nil {
		return executeError
	}

	return nil
}
