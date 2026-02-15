package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	BaseURL         string
	BlogTitle       string
	BlogDescription string
	AuthorName      string
	AdminUsername    string
	AdminPassword   string
	DBPath          string
	UploadDir       string
	SessionSecret   string
	Env             string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:            getEnv("GOBLOG_PORT", "8069"),
		BaseURL:         getEnv("GOBLOG_BASE_URL", "http://localhost:8069"),
		BlogTitle:       getEnv("GOBLOG_BLOG_TITLE", "Manas's Blog"),
		BlogDescription: getEnv("GOBLOG_BLOG_DESCRIPTION", ""),
		AuthorName:      getEnv("GOBLOG_AUTHOR_NAME", "Manas"),
		AdminUsername:    getEnv("GOBLOG_ADMIN_USERNAME", ""),
		AdminPassword:   getEnv("GOBLOG_ADMIN_PASSWORD", ""),
		DBPath:          getEnv("GOBLOG_DB_PATH", "./goblog.db"),
		UploadDir:       getEnv("GOBLOG_UPLOAD_DIR", "./uploads"),
		SessionSecret:   getEnv("GOBLOG_SESSION_SECRET", ""),
		Env:             getEnv("GOBLOG_ENV", "development"),
	}

	if cfg.Env != "development" && cfg.Env != "production" {
		log.Fatalf("invalid GOBLOG_ENV value %q: must be \"development\" or \"production\"", cfg.Env)
	}

	if cfg.Env == "production" && (cfg.AdminUsername == "" || cfg.AdminPassword == "") {
		log.Fatalf("GOBLOG_ADMIN_USERNAME and GOBLOG_ADMIN_PASSWORD are required in production")
	}

	return &cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
