package recipe

import (
	"coffie/internal/feature/recipe/domain"
	"coffie/internal/feature/recipe/http"
)

// ToCreateRecipeRequest converts an HTTP request to a domain service request.
func ToCreateRecipeRequest(req *http.CreateRecipe) domain.CreateRecipeRequest {
	return domain.CreateRecipeRequest{}
}

// ToUpdateRecipeRequest converts an HTTP update request to a domain service request.
func ToUpdateRecipeRequest(req *http.UpdateRecipe) domain.UpdateRecipeRequest {
	return domain.UpdateRecipeRequest{}
}

// ToRecipeResponse converts a detailed domain recipe to an HTTP response.
func ToRecipeResponse(r *domain.RecipeWithDetails) *http.RecipeResponse {
	return &http.RecipeResponse{}
}

// ToRecipeListItemResponse converts a recipe summary to an HTTP list item response.
func ToRecipeListItemResponse(r *domain.RecipeSummary) *http.RecipeListItem {
	return &http.RecipeListItem{}
}
