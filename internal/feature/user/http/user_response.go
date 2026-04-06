package http

import "time"

// UserResponse is the response for a single user.
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserStatsResponse holds computed user statistics for the response.
type UserStatsResponse struct {
	RecipesCount   int     `json:"recipes_count"`
	AvgRatingGiven float64 `json:"avg_rating_given"`
	FavMethod      string  `json:"fav_method"`
	FavCoffeeType  string  `json:"fav_coffee_type"`
}
