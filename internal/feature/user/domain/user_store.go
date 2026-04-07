package domain

import "context"

// UserStore defines the interface for user data persistence.
type UserStore interface {
	Create(requestContext context.Context, user *User) error
}
