package config

import "os"

const (
	defaultDatabaseURL = "postgres://coffie:coffie_pass@localhost:5432/coffie_dev?sslmode=disable"
	defaultAPIPort     = "8080"
)

// Config holds all application-level configuration.
type Config struct {
	DatabaseURL string
	APIPort     string
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
	configuration := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		APIPort:     os.Getenv("API_PORT"),
	}

	if configuration.DatabaseURL == "" {
		configuration.DatabaseURL = defaultDatabaseURL
	}

	if configuration.APIPort == "" {
		configuration.APIPort = defaultAPIPort
	}

	return configuration, nil
}
