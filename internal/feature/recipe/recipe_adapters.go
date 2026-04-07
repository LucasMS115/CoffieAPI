package recipe

import (
	"coffie/internal/feature/recipe/domain"
	"coffie/internal/feature/recipe/http"
)

// ToCreateRecipeRequest converts an HTTP request to a domain service request.
func ToCreateRecipeRequest(createRecipeRequest *http.CreateRecipe) domain.CreateRecipeRequest {
	_ = createRecipeRequest
	return domain.CreateRecipeRequest{}
}

// ToUpdateRecipeRequest converts an HTTP update request to a domain service request.
func ToUpdateRecipeRequest(updateRecipeRequest *http.UpdateRecipe) domain.UpdateRecipeRequest {
	_ = updateRecipeRequest
	return domain.UpdateRecipeRequest{}
}

// ToRecipeResponse converts a detailed domain recipe to an HTTP response.
func ToRecipeResponse(recipeWithDetails *domain.RecipeWithDetails) *http.RecipeResponse {
	_ = recipeWithDetails
	return &http.RecipeResponse{}
}

// ToRecipeListItemResponse converts a recipe summary to an HTTP list item response.
func ToRecipeListItemResponse(recipeSummary *domain.RecipeSummary) *http.RecipeListItem {
	_ = recipeSummary
	return &http.RecipeListItem{}
}
