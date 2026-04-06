package store

import (
	"context"
	"coffie/internal/feature/user/domain"
	"database/sql"
)

// PostgresUserStore implements domain.UserStore using PostgreSQL.
type PostgresUserStore struct {
	db *sql.DB
}

// NewUserStore creates a new Postgres-backed user store.
func NewUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

func (s *PostgresUserStore) Create(ctx context.Context, u *domain.User) error {
	const query = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	if _, err := s.db.ExecContext(ctx, query, u.ID, u.Name, u.Email, u.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserStore) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (s *PostgresUserStore) GetStats(ctx context.Context, userID string) (*domain.UserStats, error) {
	return nil, nil
}
