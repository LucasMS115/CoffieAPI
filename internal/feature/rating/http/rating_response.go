package http

import "time"

// RatingResponse is the response for a single rating.
type RatingResponse struct {
	ID        string     `json:"id"`
	User      UserBrief  `json:"user"`
	Score     int        `json:"score"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
}

// UserBrief holds minimal user info in nested responses.
type UserBrief struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
