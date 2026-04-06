package config

// Config holds all application-level configuration.
type Config struct {
	DatabaseURL string
	APIPort     string
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
	// load from env
	return Config{}, nil
}
