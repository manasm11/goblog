package config

import (
	"os"
	"testing"
)

// allEnvKeys lists every GOBLOG_* variable used by Load().
var allEnvKeys = []string{
	"GOBLOG_PORT",
	"GOBLOG_BASE_URL",
	"GOBLOG_BLOG_TITLE",
	"GOBLOG_BLOG_DESCRIPTION",
	"GOBLOG_AUTHOR_NAME",
	"GOBLOG_ADMIN_USERNAME",
	"GOBLOG_ADMIN_PASSWORD",
	"GOBLOG_DB_PATH",
	"GOBLOG_UPLOAD_DIR",
	"GOBLOG_SESSION_SECRET",
	"GOBLOG_ENV",
}

func clearEnv(t *testing.T) {
	t.Helper()
	for _, key := range allEnvKeys {
		os.Unsetenv(key)
	}
}

func TestLoadDefaults(t *testing.T) {
	clearEnv(t)

	cfg := Load()

	if cfg.Port != "8069" {
		t.Errorf("Port = %q, want %q", cfg.Port, "8069")
	}
	if cfg.BaseURL != "http://localhost:8069" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "http://localhost:8069")
	}
	if cfg.BlogTitle != "Manas's Blog" {
		t.Errorf("BlogTitle = %q, want %q", cfg.BlogTitle, "Manas's Blog")
	}
	if cfg.BlogDescription != "" {
		t.Errorf("BlogDescription = %q, want %q", cfg.BlogDescription, "")
	}
	if cfg.AuthorName != "Manas" {
		t.Errorf("AuthorName = %q, want %q", cfg.AuthorName, "Manas")
	}
	if cfg.AdminUsername != "" {
		t.Errorf("AdminUsername = %q, want %q", cfg.AdminUsername, "")
	}
	if cfg.AdminPassword != "" {
		t.Errorf("AdminPassword = %q, want %q", cfg.AdminPassword, "")
	}
	if cfg.DBPath != "./goblog.db" {
		t.Errorf("DBPath = %q, want %q", cfg.DBPath, "./goblog.db")
	}
	if cfg.UploadDir != "./uploads" {
		t.Errorf("UploadDir = %q, want %q", cfg.UploadDir, "./uploads")
	}
	if cfg.SessionSecret != "" {
		t.Errorf("SessionSecret = %q, want %q", cfg.SessionSecret, "")
	}
	if cfg.Env != "development" {
		t.Errorf("Env = %q, want %q", cfg.Env, "development")
	}
}

func TestLoadFromEnvVars(t *testing.T) {
	t.Setenv("GOBLOG_PORT", "9090")
	t.Setenv("GOBLOG_BASE_URL", "https://example.com")
	t.Setenv("GOBLOG_BLOG_TITLE", "Test Blog")
	t.Setenv("GOBLOG_BLOG_DESCRIPTION", "A test blog")
	t.Setenv("GOBLOG_AUTHOR_NAME", "Tester")
	t.Setenv("GOBLOG_ADMIN_USERNAME", "admin")
	t.Setenv("GOBLOG_ADMIN_PASSWORD", "secret")
	t.Setenv("GOBLOG_DB_PATH", "/tmp/test.db")
	t.Setenv("GOBLOG_UPLOAD_DIR", "/tmp/uploads")
	t.Setenv("GOBLOG_SESSION_SECRET", "mysecret")
	t.Setenv("GOBLOG_ENV", "production")

	cfg := Load()

	if cfg.Port != "9090" {
		t.Errorf("Port = %q, want %q", cfg.Port, "9090")
	}
	if cfg.BaseURL != "https://example.com" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "https://example.com")
	}
	if cfg.BlogTitle != "Test Blog" {
		t.Errorf("BlogTitle = %q, want %q", cfg.BlogTitle, "Test Blog")
	}
	if cfg.BlogDescription != "A test blog" {
		t.Errorf("BlogDescription = %q, want %q", cfg.BlogDescription, "A test blog")
	}
	if cfg.AuthorName != "Tester" {
		t.Errorf("AuthorName = %q, want %q", cfg.AuthorName, "Tester")
	}
	if cfg.AdminUsername != "admin" {
		t.Errorf("AdminUsername = %q, want %q", cfg.AdminUsername, "admin")
	}
	if cfg.AdminPassword != "secret" {
		t.Errorf("AdminPassword = %q, want %q", cfg.AdminPassword, "secret")
	}
	if cfg.DBPath != "/tmp/test.db" {
		t.Errorf("DBPath = %q, want %q", cfg.DBPath, "/tmp/test.db")
	}
	if cfg.UploadDir != "/tmp/uploads" {
		t.Errorf("UploadDir = %q, want %q", cfg.UploadDir, "/tmp/uploads")
	}
	if cfg.SessionSecret != "mysecret" {
		t.Errorf("SessionSecret = %q, want %q", cfg.SessionSecret, "mysecret")
	}
	if cfg.Env != "production" {
		t.Errorf("Env = %q, want %q", cfg.Env, "production")
	}
}

func TestLoadPartialEnvVars(t *testing.T) {
	clearEnv(t)
	t.Setenv("GOBLOG_PORT", "3000")
	t.Setenv("GOBLOG_BLOG_TITLE", "Custom Blog")

	cfg := Load()

	if cfg.Port != "3000" {
		t.Errorf("Port = %q, want %q", cfg.Port, "3000")
	}
	if cfg.BlogTitle != "Custom Blog" {
		t.Errorf("BlogTitle = %q, want %q", cfg.BlogTitle, "Custom Blog")
	}
	if cfg.BaseURL != "http://localhost:8069" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "http://localhost:8069")
	}
	if cfg.Env != "development" {
		t.Errorf("Env = %q, want %q", cfg.Env, "development")
	}
	if cfg.AuthorName != "Manas" {
		t.Errorf("AuthorName = %q, want %q", cfg.AuthorName, "Manas")
	}
	if cfg.DBPath != "./goblog.db" {
		t.Errorf("DBPath = %q, want %q", cfg.DBPath, "./goblog.db")
	}
	if cfg.UploadDir != "./uploads" {
		t.Errorf("UploadDir = %q, want %q", cfg.UploadDir, "./uploads")
	}
}

func TestGetEnv(t *testing.T) {
	t.Run("returns env var value when set", func(t *testing.T) {
		t.Setenv("TEST_CONFIG_KEY", "myvalue")
		if got := getEnv("TEST_CONFIG_KEY", "default"); got != "myvalue" {
			t.Errorf("getEnv() = %q, want %q", got, "myvalue")
		}
	})

	t.Run("returns fallback when not set", func(t *testing.T) {
		os.Unsetenv("TEST_CONFIG_KEY_MISSING")
		if got := getEnv("TEST_CONFIG_KEY_MISSING", "fallback"); got != "fallback" {
			t.Errorf("getEnv() = %q, want %q", got, "fallback")
		}
	})

	t.Run("returns fallback when set to empty", func(t *testing.T) {
		t.Setenv("TEST_CONFIG_KEY_EMPTY", "")
		if got := getEnv("TEST_CONFIG_KEY_EMPTY", "default"); got != "default" {
			t.Errorf("getEnv() = %q, want %q", got, "default")
		}
	})
}
