package user

import (
	"context"
	"errors"
	"testing"

	"coffie/internal/feature/user/domain"
)

// --- success: register user ---
func TestService_Register_Success(testingContext *testing.T) {
	// given
	mockStore := &mockUserStore{
		CreateFunction: func(requestContext context.Context, user *domain.User) error {
			_ = requestContext
			expected := &domain.User{
				Name:  "Lucas",
				Email: "lucas@email.com",
			}
			if user.Name != expected.Name {
				testingContext.Errorf("expected name %q, got %q", expected.Name, user.Name)
			}
			if user.Email != expected.Email {
				testingContext.Errorf("expected email %q, got %q", expected.Email, user.Email)
			}
			if user.ID == "" {
				testingContext.Error("expected ID to be generated")
			}
			if user.CreatedAt.IsZero() {
				testingContext.Error("expected CreatedAt to be set")
			}
			return nil
		},
	}

	userService := domain.NewService(mockStore)
	registerRequest := domain.RegisterRequest{
		Name:  "Lucas",
		Email: "lucas@email.com",
	}

	// when
	result, registerError := userService.Register(context.Background(), registerRequest)

	// then
	if registerError != nil {
		testingContext.Fatalf("expected no error, got: %v", registerError)
	}
	if result == nil {
		testingContext.Fatal("expected result, got nil")
	}
}

// --- failure: register user passes through store error ---
func TestService_Register_Failure_StoreError(testingContext *testing.T) {
	// given
	mockStore := &mockUserStore{
		CreateFunction: func(requestContext context.Context, user *domain.User) error {
			_ = requestContext
			_ = user
			return errors.New("db connection refused")
		},
	}

	userService := domain.NewService(mockStore)
	registerRequest := domain.RegisterRequest{
		Name:  "Lucas",
		Email: "lucas@email.com",
	}

	// when
	_, registerError := userService.Register(context.Background(), registerRequest)

	// then
	if registerError == nil {
		testingContext.Fatal("expected error, got nil")
	}
	if registerError.Error() != "db connection refused" {
		testingContext.Errorf("expected 'db connection refused', got: %v", registerError)
	}
}
