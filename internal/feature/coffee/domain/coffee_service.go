package domain

import "context"

// CoffeeStore defines the interface for coffee data persistence.
type CoffeeStore interface {
	Create(requestContext context.Context, coffee *Coffee) error
	GetByID(requestContext context.Context, coffeeID string) (*Coffee, error)
	List(requestContext context.Context, filter ListFilter) ([]Coffee, int, error)
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
func (service *Service) Create(requestContext context.Context, createCoffeeRequest CreateCoffeeRequest) (*Coffee, error) {
	_ = requestContext
	_ = createCoffeeRequest
	return nil, nil
}

// GetByID returns a coffee by ID.
func (service *Service) GetByID(requestContext context.Context, coffeeID string) (*Coffee, error) {
	_ = service
	_ = requestContext
	_ = coffeeID
	return nil, nil
}

// List searches and lists coffees.
func (service *Service) List(requestContext context.Context, listCoffeesRequest ListCoffeesRequest) ([]Coffee, int, error) {
	_ = service
	_ = requestContext
	_ = listCoffeesRequest
	return nil, 0, nil
}
