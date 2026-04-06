package http

import "time"

// CoffeeResponse is the response for a single coffee.
type CoffeeResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Type        string    `json:"type"`
	FlavorNotes string    `json:"flavor_notes"`
	CreatedAt   time.Time `json:"created_at"`
}

// CoffeeListItem is a summarized coffee for list responses.
type CoffeeListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}

// ListEnvelope wraps list responses with pagination.
type ListEnvelope struct {
	Items []CoffeeListItem `json:"items"`
	Total int              `json:"total"`
	Page  int              `json:"page"`
}
