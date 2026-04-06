package http

import (
	"database/sql"
	"net/http"

	"coffie/internal/feature/user/domain"
	userhttp "coffie/internal/feature/user/http"
	"coffie/internal/feature/user/store"
)

// NewServer creates and configures the HTTP server with all routes.
func NewServer(addr string, db *sql.DB) *http.Server {
	mux := http.NewServeMux()

	healthHandler := NewHealthHandler()
	healthHandler.RegisterRoutes(mux)

	userStore := store.NewUserStore(db)
	userSvc := domain.NewService(userStore)
	userHandler := userhttp.NewHandler(userSvc)
	userHandler.RegisterRoutes(mux)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
