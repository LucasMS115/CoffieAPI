package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
)

// --- mock store ---
type mockUserStore struct {
	CreateFn  func(ctx context.Context, u *domain.User) error
	GetByIDFn func(ctx context.Context, id string) (*domain.User, error)
}

func (m *mockUserStore) Create(ctx context.Context, u *domain.User) error {
	return m.CreateFn(ctx, u)
}

func (m *mockUserStore) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return m.GetByIDFn(ctx, id)
}

func (m *mockUserStore) GetStats(ctx context.Context, userID string) (*domain.UserStats, error) {
	return nil, nil
}

// --- success: POST /api/users ---
func TestRegisterUser_Success(t *testing.T) {
	// given
	svc := domain.NewService(&mockUserStore{
		CreateFn: func(ctx context.Context, u *domain.User) error { return nil },
	})

	handler := userhttp.NewHandler(svc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	body := `{"name":"Lucas","email":"lucas@email.com"}`
	req := httptest.NewRequest("POST", "/api/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// when
	mux.ServeHTTP(rec, req)

	// then
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var resp userhttp.UserResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("expected valid JSON, got error: %v", err)
	}

	if resp.Name != "Lucas" {
		t.Errorf("expected name 'Lucas', got %q", resp.Name)
	}
	if resp.Email != "lucas@email.com" {
		t.Errorf("expected email 'lucas@email.com', got %q", resp.Email)
	}
	if resp.ID == "" {
		t.Error("expected id to be non-empty")
	}
}
