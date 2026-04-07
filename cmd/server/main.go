package main

import (
	"fmt"
	"log"
	"net/http"

	_ "coffie/docs"

	_ "github.com/lib/pq"

	"coffie/internal/config"
	"coffie/internal/database"
	apphttp "coffie/internal/http"
)

// @title           Coffee API
// @version         1.0
// @description     REST API for managing the coffie app.
// @host            localhost:8080
// @BasePath        /
// @schemes         http
func main() {
	applicationConfig, loadConfigError := config.Load()
	if loadConfigError != nil {
		log.Fatalf("failed to load configuration: %v", loadConfigError)
	}

	databaseConnection, openDatabaseError := database.NewPostgresConn(applicationConfig.DatabaseURL)
	if openDatabaseError != nil {
		log.Fatalf("failed to open database: %v", openDatabaseError)
	}
	defer databaseConnection.Close()

	fmt.Printf("connected to database\n")

	server := apphttp.NewServer(":"+applicationConfig.APIPort, databaseConnection)

	fmt.Printf("server starting on port %s\n", applicationConfig.APIPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
