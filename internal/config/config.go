package config

import (
	"log"
	"os"
)

type Config struct {
	Addr   string
	DbURL  string
}

func Load() Config {
	addr := getEnv("ADDR", ":8080")
	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	return Config{Addr: addr, DbURL: dbURL}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
