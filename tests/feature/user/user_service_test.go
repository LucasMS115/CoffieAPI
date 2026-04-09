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
	expectedName := "Lucas"
	expectedEmail := "lucas@email.com"
	mockStore := &mockUserStore{
		CreateFunction: func(requestContext context.Context, user *domain.User) error {
			_ = requestContext
			assertRegisteredUser(testingContext, user, expectedName, expectedEmail)
			return nil
		},
	}

	userService := domain.NewService(mockStore)
	registerRequest := domain.RegisterRequest{
		Name:  expectedName,
		Email: expectedEmail,
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
	assertRegisteredUser(testingContext, result, expectedName, expectedEmail)
}

// --- failure: register user passes through store error ---
func TestService_Register_Failure_StoreError(testingContext *testing.T) {
	// given
	expectedError := errors.New("db connection refused")
	mockStore := &mockUserStore{
		CreateFunction: func(requestContext context.Context, user *domain.User) error {
			_ = requestContext
			_ = user
			return expectedError
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
	if !errors.Is(registerError, expectedError) {
		testingContext.Errorf("expected propagated error %v, got: %v", expectedError, registerError)
	}
}

func assertRegisteredUser(testingContext *testing.T, user *domain.User, expectedName string, expectedEmail string) {
	testingContext.Helper()

	if user == nil {
		testingContext.Fatal("expected user, got nil")
	}
	if user.Name != expectedName {
		testingContext.Errorf("expected name %q, got %q", expectedName, user.Name)
	}
	if user.Email != expectedEmail {
		testingContext.Errorf("expected email %q, got %q", expectedEmail, user.Email)
	}
	if user.ID == "" {
		testingContext.Error("expected ID to be generated")
	}
	if user.CreatedAt.IsZero() {
		testingContext.Error("expected CreatedAt to be set")
	}
}
