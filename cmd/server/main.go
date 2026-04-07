package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "coffie/docs"

	_ "github.com/lib/pq"

	apphttp "coffie/internal/http"
)

// @title           Coffee API
// @version         1.0
// @description     REST API for managing the coffie app.
// @host            localhost:8080
// @BasePath        /
// @schemes         http
func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://coffie:coffie_pass@localhost:5432/coffie_dev?sslmode=disable"
	}

	databaseConnection, openDatabaseError := openDatabase(databaseURL)
	if openDatabaseError != nil {
		log.Fatalf("failed to open database: %v", openDatabaseError)
	}
	defer databaseConnection.Close()

	fmt.Printf("connected to database\n")

	server := apphttp.NewServer(":"+port, databaseConnection)

	fmt.Printf("server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func openDatabase(dataSourceName string) (*sql.DB, error) {
	databaseConnection, openError := sql.Open("postgres", dataSourceName)
	if openError != nil {
		return nil, openError
	}
	if pingError := databaseConnection.Ping(); pingError != nil {
		return nil, pingError
	}
	return databaseConnection, nil
}
