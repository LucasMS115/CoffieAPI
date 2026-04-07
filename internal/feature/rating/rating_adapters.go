package rating

import (
	"coffie/internal/feature/rating/domain"
	"coffie/internal/feature/rating/http"
)

// ToCreateRatingRequest converts an HTTP request to a domain service request.
func ToCreateRatingRequest(createRatingRequest *http.CreateRating) domain.CreateRatingRequest {
	_ = createRatingRequest
	return domain.CreateRatingRequest{}
}

// ToRatingResponse converts a domain rating to an HTTP response.
func ToRatingResponse(rating *domain.Rating) *http.RatingResponse {
	_ = rating
	return &http.RatingResponse{}
}
