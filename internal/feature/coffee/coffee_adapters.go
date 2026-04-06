package coffee

import (
	"coffie/internal/feature/coffee/domain"
	"coffie/internal/feature/coffee/http"
)

// ToCreateCoffeeRequest converts an HTTP request to a domain service request.
func ToCreateCoffeeRequest(req *http.CreateCoffee) domain.CreateCoffeeRequest {
	return domain.CreateCoffeeRequest{}
}

// ToCoffeeResponse converts a domain Coffee to an HTTP response.
func ToCoffeeResponse(c *domain.Coffee) *http.CoffeeResponse {
	return &http.CoffeeResponse{}
}
