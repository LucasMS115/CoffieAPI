package domain

import "context"

// CoffeeStore defines the interface for coffee data persistence.
type CoffeeStore interface {
	Create(ctx context.Context, c *Coffee) error
	GetByID(ctx context.Context, id string) (*Coffee, error)
	List(ctx context.Context, filter ListFilter) ([]Coffee, int, error)
}

// ListFilter holds optional filters for listing coffees.
type ListFilter struct {
	Search string
	Type   string
	Page   int
	Limit  int
}

// Service handles coffee business logic.
type Service struct {
	coffeeStore CoffeeStore
}

// NewService creates a new Service.
func NewService(coffeeStore CoffeeStore) *Service {
	return &Service{coffeeStore: coffeeStore}
}

// CreateCoffeeRequest holds validated input for creating a coffee.
type CreateCoffeeRequest struct {
	Name        string
	Brand       string
	Type        string
	FlavorNotes string
}

// ListCoffeesRequest holds validated query params for listing coffees.
type ListCoffeesRequest struct {
	Search string
	Type   string
	Page   int
	Limit  int
}

// Create registers a new coffee.
func (s *Service) Create(ctx context.Context, req CreateCoffeeRequest) (*Coffee, error) {
	return nil, nil
}

// GetByID returns a coffee by ID.
func (s *Service) GetByID(ctx context.Context, id string) (*Coffee, error) {
	return nil, nil
}

// List searches and lists coffees.
func (s *Service) List(ctx context.Context, req ListCoffeesRequest) ([]Coffee, int, error) {
	return nil, 0, nil
}
