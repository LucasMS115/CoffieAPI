package http

// CreateCoffee holds the validated body for POST /api/coffees.
type CreateCoffee struct {
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Type        string `json:"type"`
	FlavorNotes string `json:"flavor_notes"`
}
