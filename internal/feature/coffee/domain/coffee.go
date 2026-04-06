package domain

import "time"

// Coffee represents a coffee bean/product in the catalog.
type Coffee struct {
	ID          string
	Name        string
	Brand       string
	Type        string
	FlavorNotes string
	CreatedAt   time.Time
}
