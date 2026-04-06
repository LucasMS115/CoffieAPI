package domain

import "time"

// Rating represents a user's score on a recipe.
type Rating struct {
	ID        string
	RecipeID  string
	UserID    string
	Score     int
	Comment   string
	CreatedAt time.Time
}
