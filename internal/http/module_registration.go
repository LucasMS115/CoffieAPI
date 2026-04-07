package http

import (
	"database/sql"
	"net/http"

	userfeature "coffie/internal/feature/user"

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
	userModule := userfeature.NewModule(databaseConnection)
	userModule.RegisterRoutes(serveMux)
}
