package http

// RegisterUser holds the validated body for POST /api/users.
type RegisterUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
