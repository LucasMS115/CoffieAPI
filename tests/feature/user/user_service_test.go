package user

import (
	"context"
	"errors"
	"testing"

	"coffie/internal/feature/user/domain"
)

// --- given ---
type mockUserStore struct {
	CreateFn   func(ctx context.Context, u *domain.User) error
	GetByIDFn  func(ctx context.Context, id string) (*domain.User, error)
	GetStatsFn func(ctx context.Context, userID string) (*domain.UserStats, error)
}

func (m *mockUserStore) Create(ctx context.Context, u *domain.User) error {
	return m.CreateFn(ctx, u)
}

func (m *mockUserStore) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return m.GetByIDFn(ctx, id)
}

func (m *mockUserStore) GetStats(ctx context.Context, userID string) (*domain.UserStats, error) {
	return m.GetStatsFn(ctx, userID)
}

// --- success: register user ---
func TestService_Register_Success(t *testing.T) {
	// given
	store := &mockUserStore{
		CreateFn: func(ctx context.Context, u *domain.User) error {
			expected := &domain.User{
				Name:  "Lucas",
				Email: "lucas@email.com",
			}
			if u.Name != expected.Name {
				t.Errorf("expected name %q, got %q", expected.Name, u.Name)
			}
			if u.Email != expected.Email {
				t.Errorf("expected email %q, got %q", expected.Email, u.Email)
			}
			if u.ID == "" {
				t.Error("expected ID to be generated")
			}
			if u.CreatedAt.IsZero() {
				t.Error("expected CreatedAt to be set")
			}
			return nil
		},
	}

	svc := domain.NewService(store)
	req := domain.RegisterRequest{
		Name:  "Lucas",
		Email: "lucas@email.com",
	}

	// when
	result, err := svc.Register(context.Background(), req)

	// then
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

// --- failure: register user passes through store error ---
func TestService_Register_Failure_StoreError(t *testing.T) {
	// given
	store := &mockUserStore{
		CreateFn: func(ctx context.Context, u *domain.User) error {
			return errors.New("db connection refused")
		},
	}

	svc := domain.NewService(store)
	req := domain.RegisterRequest{
		Name:  "Lucas",
		Email: "lucas@email.com",
	}

	// when
	_, err := svc.Register(context.Background(), req)

	// then
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "db connection refused" {
		t.Errorf("expected 'db connection refused', got: %v", err)
	}
}
