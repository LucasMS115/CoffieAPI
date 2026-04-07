package domain

import "context"

// UserStore defines the interface for user data persistence.
type UserStore interface {
	Create(ctx context.Context, user *User) error
}
