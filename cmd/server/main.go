package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "coffie/docs"

	_ "github.com/lib/pq"
	_ "github.com/swaggo/http-swagger"

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

	db, err := openDB(databaseURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	fmt.Printf("connected to database\n")

	server := apphttp.NewServer(":"+port, db)

	fmt.Printf("server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
