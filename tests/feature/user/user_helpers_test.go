package user

import (
	"context"
	"net/http"

	"coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
)

type mockUserStore struct {
	CreateFunction func(requestContext context.Context, user *domain.User) error
}

func (mockStore *mockUserStore) Create(requestContext context.Context, user *domain.User) error {
	if mockStore.CreateFunction == nil {
		return nil
	}
	return mockStore.CreateFunction(requestContext, user)
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
