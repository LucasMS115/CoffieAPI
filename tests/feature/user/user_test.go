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

// --- success: POST /api/users ---
func TestRegisterUser_Success(testingContext *testing.T) {
	// given
	userService := domain.NewService(&mockUserStore{
		CreateFunction: func(requestContext context.Context, user *domain.User) error {
			_ = requestContext
			_ = user
			return nil
		},
	})

	userHandler := userhttp.NewHandler(userService)

	serveMux := http.NewServeMux()
	userHandler.RegisterRoutes(serveMux)

	requestBody := `{"name":"Lucas","email":"lucas@email.com"}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// when
	serveMux.ServeHTTP(responseRecorder, request)

	// then
	if responseRecorder.Code != http.StatusCreated {
		testingContext.Fatalf("expected status %d, got %d", http.StatusCreated, responseRecorder.Code)
	}

	var userResponse userhttp.UserResponse
	if decodeError := json.NewDecoder(responseRecorder.Body).Decode(&userResponse); decodeError != nil {
		testingContext.Fatalf("expected valid JSON, got error: %v", decodeError)
	}

	if userResponse.Name != "Lucas" {
		testingContext.Errorf("expected name 'Lucas', got %q", userResponse.Name)
	}
	if userResponse.Email != "lucas@email.com" {
		testingContext.Errorf("expected email 'lucas@email.com', got %q", userResponse.Email)
	}
	if userResponse.ID == "" {
		testingContext.Error("expected id to be non-empty")
	}
}
