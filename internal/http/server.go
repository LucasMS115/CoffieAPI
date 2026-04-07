package http

import (
	"database/sql"
	"net/http"
)

// NewServer creates and configures the HTTP server with all routes.
func NewServer(address string, databaseConnection *sql.DB) *http.Server {
	serveMux := http.NewServeMux()

	registerSwaggerModule(serveMux)
	registerHealthModule(serveMux)
	registerUserModule(serveMux, databaseConnection)

	return &http.Server{
		Addr:    address,
		Handler: serveMux,
	}
}
