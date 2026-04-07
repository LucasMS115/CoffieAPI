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
func setupUserStore(testingContext *testing.T) (*store.PostgresUserStore, sqlmock.Sqlmock) {
	testingContext.Helper()

	databaseConnection, sqlMock, setupError := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if setupError != nil {
		testingContext.Fatalf("failed to create sqlmock: %v", setupError)
	}

	userStore := store.NewUserStore(databaseConnection)
	return userStore, sqlMock
}

// --- success: create user ---
func TestPostgresUserStore_Create_Success(testingContext *testing.T) {
	// given
	userStore, sqlMock := setupUserStore(testingContext)

	const insertSQL = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	user := &domain.User{
		ID:        "test-id",
		Name:      "Lucas",
		Email:     "lucas@email.com",
		CreatedAt: time.Date(2026, 4, 5, 10, 0, 0, 0, time.UTC),
	}

	sqlMock.ExpectExec(insertSQL).
		WithArgs(user.ID, user.Name, user.Email, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// when
	createError := userStore.Create(context.Background(), user)

	// then
	if createError != nil {
		testingContext.Errorf("expected no error, got: %v", createError)
	}
	if expectationsError := sqlMock.ExpectationsWereMet(); expectationsError != nil {
		testingContext.Errorf("unmet expectations: %v", expectationsError)
	}
}

// --- failure: create user returns error on database failure ---
func TestPostgresUserStore_Create_Failure(testingContext *testing.T) {
	// given
	userStore, sqlMock := setupUserStore(testingContext)

	const insertSQL = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`

	user := &domain.User{
		ID:        "test-id",
		Name:      "Lucas",
		Email:     "lucas@email.com",
		CreatedAt: time.Date(2026, 4, 5, 10, 0, 0, 0, time.UTC),
	}

	sqlMock.ExpectExec(insertSQL).
		WithArgs(user.ID, user.Name, user.Email, user.CreatedAt).
		WillReturnError(errors.New("connection refused"))

	// when
	createError := userStore.Create(context.Background(), user)

	// then
	if createError == nil {
		testingContext.Error("expected error, got nil")
	}
	if expectationsError := sqlMock.ExpectationsWereMet(); expectationsError != nil {
		testingContext.Errorf("unmet expectations: %v", expectationsError)
	}
}
