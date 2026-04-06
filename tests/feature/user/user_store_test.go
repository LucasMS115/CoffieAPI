package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"coffie/internal/feature/user/domain"
	"coffie/internal/feature/user/store"
)

// --- given ---
func setupStore(t *testing.T) (*store.PostgresUserStore, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	s := store.NewUserStore(db)
	return s, mock
}

// --- success: create user ---
func TestPostgresUserStore_Create_Success(t *testing.T) {
	// given
	s, mock := setupStore(t)

	const insertSQL = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	u := &domain.User{
		ID:        "test-id",
		Name:      "Lucas",
		Email:     "lucas@email.com",
		CreatedAt: time.Date(2026, 4, 5, 10, 0, 0, 0, time.UTC),
	}

	mock.ExpectExec(insertSQL).
		WithArgs(u.ID, u.Name, u.Email, u.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// when
	err := s.Create(context.Background(), u)

	// then
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

// --- failure: create user returns error on database failure ---
func TestPostgresUserStore_Create_Failure(t *testing.T) {
	// given
	s, mock := setupStore(t)

	const insertSQL = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	u := &domain.User{
		ID:        "test-id",
		Name:      "Lucas",
		Email:     "lucas@email.com",
		CreatedAt: time.Date(2026, 4, 5, 10, 0, 0, 0, time.UTC),
	}

	mock.ExpectExec(insertSQL).
		WithArgs(u.ID, u.Name, u.Email, u.CreatedAt).
		WillReturnError(errors.New("connection refused"))

	// when
	err := s.Create(context.Background(), u)

	// then
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
