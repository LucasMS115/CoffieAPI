package user

import (
	"context"

	"coffie/internal/feature/user/domain"
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
