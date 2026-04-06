package http

// CreateRecipe holds the validated body for POST /api/recipes.
type CreateRecipe struct {
	CoffeeID    string  `json:"coffee_id"`
	Method      string  `json:"method"`
	WaterTemp   int     `json:"water_temp"`
	Dose        float64 `json:"dose"`
	Yield       float64 `json:"yield"`
	BrewTime    int     `json:"brew_time"`
	Description string  `json:"description"`
}

// UpdateRecipe holds the validated body for PUT /api/recipes/{id}.
type UpdateRecipe struct {
	Method      *string  `json:"method,omitempty"`
	WaterTemp   *int     `json:"water_temp,omitempty"`
	Dose        *float64 `json:"dose,omitempty"`
	Yield       *float64 `json:"yield,omitempty"`
	BrewTime    *int     `json:"brew_time,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// ListRecipes holds validated query params for GET /api/recipes.
type ListRecipes struct {
	UserID string
	Method string
	Search string
	Page   int
	Limit  int
}
