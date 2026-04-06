package http

// CreateRating holds the validated body for POST /api/recipes/{id}/ratings.
type CreateRating struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}
