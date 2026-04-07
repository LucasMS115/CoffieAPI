package coffee

import (
	"coffie/internal/feature/coffee/domain"
	"coffie/internal/feature/coffee/http"
)

// ToCreateCoffeeRequest converts an HTTP request to a domain service request.
func ToCreateCoffeeRequest(createCoffeeRequest *http.CreateCoffee) domain.CreateCoffeeRequest {
	_ = createCoffeeRequest
	return domain.CreateCoffeeRequest{}
}

// ToCoffeeResponse converts a domain Coffee to an HTTP response.
func ToCoffeeResponse(coffee *domain.Coffee) *http.CoffeeResponse {
	_ = coffee
	return &http.CoffeeResponse{}
}
