package http

import "time"

// RecipeResponse is the full recipe detail response.
type RecipeResponse struct {
	ID          string      `json:"id"`
	User        UserBrief   `json:"user"`
	Coffee      CoffeeBrief `json:"coffee"`
	Method      string      `json:"method"`
	WaterTemp   int         `json:"water_temp"`
	Dose        float64     `json:"dose"`
	Yield       float64     `json:"yield"`
	BrewTime    int         `json:"brew_time"`
	Description string      `json:"description"`
	AvgRating   *float64    `json:"avg_rating"`
	RatingCount int         `json:"rating_count"`
	CreatedAt   time.Time   `json:"created_at"`
}

// RecipeListItem is a summarized recipe for list views.
type RecipeListItem struct {
	ID          string      `json:"id"`
	User        UserBrief   `json:"user"`
	Coffee      CoffeeBrief `json:"coffee"`
	Method      string      `json:"method"`
	AvgRating   *float64    `json:"avg_rating"`
	RatingCount int         `json:"rating_count"`
	CreatedAt   time.Time   `json:"created_at"`
}

// ListEnvelope wraps list responses with pagination.
type ListEnvelope struct {
	Items []RecipeListItem `json:"items"`
	Total int              `json:"total"`
	Page  int              `json:"page"`
}

// UserBrief holds minimal user info in nested responses.
type UserBrief struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CoffeeBrief holds minimal coffee info in nested responses.
type CoffeeBrief struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}
