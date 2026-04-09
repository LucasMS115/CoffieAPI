package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
	"coffie/internal/http/response"
)

var registerUserCreatedAt = time.Date(2026, 4, 9, 13, 0, 0, 0, time.UTC)

func TestRegisterUser_Success(testingContext *testing.T) {
	userService := &mockUserService{
		RegisterFunction: func(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error) {
			_ = requestContext

			if registerRequest.Name != "Lucas" {
				testingContext.Fatalf("expected name 'Lucas', got %q", registerRequest.Name)
			}
			if registerRequest.Email != "lucas@email.com" {
				testingContext.Fatalf("expected email 'lucas@email.com', got %q", registerRequest.Email)
			}

			return &domain.User{
				ID:        "generated-id",
				Name:      registerRequest.Name,
				Email:     registerRequest.Email,
				CreatedAt: registerUserCreatedAt,
			}, nil
		},
	}

	serveMux := newUserServeMux(userService)

	requestBody := `{"name":"Lucas","email":"lucas@email.com"}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	serveMux.ServeHTTP(responseRecorder, request)

	assertRegisterUserResponse(testingContext, responseRecorder, http.StatusCreated, userhttp.UserResponse{
		ID:        "generated-id",
		Name:      "Lucas",
		Email:     "lucas@email.com",
		CreatedAt: registerUserCreatedAt,
	})
}

func TestRegisterUser_InvalidJSON(testingContext *testing.T) {
	serviceCalled := false
	serveMux := newUserServeMux(&mockUserService{
		RegisterFunction: func(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error) {
			_ = requestContext
			_ = registerRequest
			serviceCalled = true
			return nil, nil
		},
	})

	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(`{"name":`))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	serveMux.ServeHTTP(responseRecorder, request)

	if serviceCalled {
		testingContext.Fatal("expected service not to be called")
	}
	assertRegisterErrorResponse(testingContext, responseRecorder, http.StatusBadRequest, response.ErrorResponse{
		ErrorCode: "INVALID_INPUT",
		Message:   "invalid request body",
	})
}

func TestRegisterUser_ValidationFailure(testingContext *testing.T) {
	serviceCalled := false
	serveMux := newUserServeMux(&mockUserService{
		RegisterFunction: func(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error) {
			_ = requestContext
			_ = registerRequest
			serviceCalled = true
			return nil, nil
		},
	})

	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(`{"name":"Lucas"}`))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	serveMux.ServeHTTP(responseRecorder, request)

	if serviceCalled {
		testingContext.Fatal("expected service not to be called")
	}
	assertRegisterErrorResponse(testingContext, responseRecorder, http.StatusBadRequest, response.ErrorResponse{
		ErrorCode: "INVALID_INPUT",
		Message:   "validation failed",
		Fields: []response.FieldError{
			{Field: "email", Message: "is required"},
		},
	})
}

func TestRegisterUser_ServiceError(testingContext *testing.T) {
	serveMux := newUserServeMux(&mockUserService{
		RegisterFunction: func(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error) {
			_ = requestContext
			_ = registerRequest
			return nil, domain.ErrUserAlreadyExists
		},
	})

	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(`{"name":"Lucas","email":"lucas@email.com"}`))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	serveMux.ServeHTTP(responseRecorder, request)

	assertRegisterErrorResponse(testingContext, responseRecorder, http.StatusConflict, response.ErrorResponse{
		ErrorCode: "CONFLICT",
		Message:   "user already exists",
	})
}

type mockUserService struct {
	RegisterFunction func(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error)
}

func (mockService *mockUserService) Register(requestContext context.Context, registerRequest domain.RegisterRequest) (*domain.User, error) {
	if mockService.RegisterFunction == nil {
		return nil, nil
	}
	return mockService.RegisterFunction(requestContext, registerRequest)
}

func newUserServeMux(userService userhttp.UserService) *http.ServeMux {
	userHandler := userhttp.NewHandler(userService)
	serveMux := http.NewServeMux()
	userHandler.RegisterRoutes(serveMux)
	return serveMux
}

func assertRegisterErrorResponse(testingContext *testing.T, responseRecorder *httptest.ResponseRecorder, expectedStatus int, expectedResponse response.ErrorResponse) {
	testingContext.Helper()

	if responseRecorder.Code != expectedStatus {
		testingContext.Fatalf("expected status %d, got %d", expectedStatus, responseRecorder.Code)
	}
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		testingContext.Fatalf("expected Content-Type application/json, got %q", responseRecorder.Header().Get("Content-Type"))
	}

	var errorResponse response.ErrorResponse
	if decodeError := json.NewDecoder(responseRecorder.Body).Decode(&errorResponse); decodeError != nil {
		testingContext.Fatalf("expected valid JSON error response, got error: %v", decodeError)
	}

	if !reflect.DeepEqual(errorResponse, expectedResponse) {
		testingContext.Fatalf("expected error response %#v, got %#v", expectedResponse, errorResponse)
	}
}

func assertRegisterUserResponse(testingContext *testing.T, responseRecorder *httptest.ResponseRecorder, expectedStatus int, expectedResponse userhttp.UserResponse) {
	testingContext.Helper()

	if responseRecorder.Code != expectedStatus {
		testingContext.Fatalf("expected status %d, got %d", expectedStatus, responseRecorder.Code)
	}
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		testingContext.Fatalf("expected Content-Type application/json, got %q", responseRecorder.Header().Get("Content-Type"))
	}

	var userResponse userhttp.UserResponse
	if decodeError := json.NewDecoder(responseRecorder.Body).Decode(&userResponse); decodeError != nil {
		testingContext.Fatalf("expected valid JSON user response, got error: %v", decodeError)
	}

	if !reflect.DeepEqual(userResponse, expectedResponse) {
		testingContext.Fatalf("expected user response %#v, got %#v", expectedResponse, userResponse)
	}
}
