package domain

import "time"

// User represents a registered user in the system.
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}
