package user

import (
	"database/sql"
	"net/http"

	userdomain "coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
	userstore "coffie/internal/feature/user/store"
)

// Module groups the user feature dependencies and route registration.
type Module struct {
	userHandler *userhttp.Handler
}

// NewModule creates the user feature module.
func NewModule(databaseConnection *sql.DB) *Module {
	userStore := userstore.NewUserStore(databaseConnection)
	userService := userdomain.NewService(userStore)
	userHandler := userhttp.NewHandler(userService)

	return &Module{userHandler: userHandler}
}

// RegisterRoutes attaches the user feature routes to the provided ServeMux.
func (module *Module) RegisterRoutes(serveMux *http.ServeMux) {
	module.userHandler.RegisterRoutes(serveMux)
}
