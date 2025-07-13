package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL   string
	SupabaseKey   string
	SupabaseDBURL string
	JWTSecret     string
	Port          string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		SupabaseURL:   getEnv("SUPABASE_URL", ""),
		SupabaseKey:   getEnv("SUPABASE_KEY", ""),
		SupabaseDBURL: getEnv("SUPABASE_DB_URL", ""),
		JWTSecret:     getEnv("JWT_SECRET", "secret"),
		Port:          getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
