package http

import (
	"database/sql"
	"net/http"

	userdomain "coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
	userstore "coffie/internal/feature/user/store"

	httpSwagger "github.com/swaggo/http-swagger"
)

func registerSwaggerModule(serveMux *http.ServeMux) {
	serveMux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
}

func registerHealthModule(serveMux *http.ServeMux) {
	healthHandler := NewHealthHandler()
	healthHandler.RegisterRoutes(serveMux)
}

func registerUserModule(serveMux *http.ServeMux, databaseConnection *sql.DB) {
	userStore := userstore.NewUserStore(databaseConnection)
	userService := userdomain.NewService(userStore)
	userHandler := userhttp.NewHandler(userService)
	userHandler.RegisterRoutes(serveMux)
}
