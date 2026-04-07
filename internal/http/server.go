package http

import (
	"database/sql"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
	"coffie/internal/feature/user/store"
)

// NewServer creates and configures the HTTP server with all routes.
func NewServer(address string, databaseConnection *sql.DB) *http.Server {
	serveMux := http.NewServeMux()

	// Swagger documentation
	serveMux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	healthHandler := NewHealthHandler()
	healthHandler.RegisterRoutes(serveMux)

	userStore := store.NewUserStore(databaseConnection)
	userService := domain.NewService(userStore)
	userHandler := userhttp.NewHandler(userService)
	userHandler.RegisterRoutes(serveMux)

	return &http.Server{
		Addr:    address,
		Handler: serveMux,
	}
}
