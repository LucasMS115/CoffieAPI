package rating

import (
	"coffie/internal/feature/rating/domain"
	"coffie/internal/feature/rating/http"
)

// ToCreateRatingRequest converts an HTTP request to a domain service request.
func ToCreateRatingRequest(req *http.CreateRating) domain.CreateRatingRequest {
	return domain.CreateRatingRequest{}
}

// ToRatingResponse converts a domain rating to an HTTP response.
func ToRatingResponse(r *domain.Rating) *http.RatingResponse {
	return &http.RatingResponse{}
}
