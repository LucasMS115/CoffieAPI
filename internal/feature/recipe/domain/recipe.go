package domain

import "time"

// Recipe represents a coffee brewing recipe.
type Recipe struct {
	ID          string
	UserID      string
	CoffeeID    string
	Method      string
	WaterTemp   int
	Dose        float64
	Yield       float64
	BrewTime    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// RecipeSummary holds a recipe with aggregated data for list views.
type RecipeSummary struct {
	Recipe
	AvgRating   *float64
	RatingCount int
}

// RecipeWithDetails has all recipe data plus nested coffee and user info from joins.
type RecipeWithDetails struct {
	ID          string
	UserID      string
	UserName    string
	CoffeeID    string
	CoffeeName  string
	CoffeeBrand string
	Method      string
	WaterTemp   int
	Dose        float64
	Yield       float64
	BrewTime    int
	Description string
	AvgRating   *float64
	RatingCount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
