# CLAUDE.md — Project Reference for Claude Code

## Project

GoBlog — a personal blogging platform.
Module path: `github.com/manasm11/goblog`
Go version: 1.25.6

## Architecture

Layered architecture with clear separation of concerns:

```
cmd/server/main.go    — entry point, HTTP server with graceful shutdown
internal/config/      — environment-based configuration
internal/database/    — SQLite setup, migrations
internal/models/      — data structs
internal/repository/  — data access (CRUD, queries)
internal/services/    — business logic
internal/handlers/    — HTTP handlers
internal/middleware/   — HTTP middleware
internal/markdown/    — markdown rendering
internal/seo/         — sitemap, RSS, structured data
templates/            — Templ templates (layouts/, pages/, partials/, admin/)
static/               — static assets (css/, js/, images/)
uploads/              — user file uploads
```

## Key Conventions

- **Port**: 8069 (configured via `GOBLOG_PORT`)
- **Routing**: Go 1.22+ `http.ServeMux` with method patterns (e.g., `"GET /health"`)
- **Graceful shutdown**: `signal.NotifyContext` with SIGINT/SIGTERM, 10s timeout
- **Config**: `internal/config.Load()` reads `.env` via godotenv, then env vars with `GOBLOG_` prefix, then defaults. In production, `GOBLOG_ADMIN_USERNAME` and `GOBLOG_ADMIN_PASSWORD` are required (log.Fatal).
- **Empty dirs**: Tracked via `.gitkeep` files
- **Templates**: Templ (type-safe Go HTML templating)
- **CSS framework**: Pico CSS (classless)
- **Database**: SQLite (file: `goblog.db`, gitignored)

## Build & Run

```bash
make run     # go run cmd/server/main.go
make build   # go build -o bin/goblog cmd/server/main.go
make test    # go test ./... -v
make templ   # templ generate
make dev     # templ generate && go run cmd/server/main.go
make clean   # rm -rf bin/ tmp/
```

## Health Endpoint

`GET /health` returns `{"status":"ok","version":"0.1.0"}` with `Content-Type: application/json`.

## Gitignore Rules

Ignored: `goblog.db`, `.env`, `uploads/*` (except `.gitkeep`), `bin/`, `tmp/`, `*.exe`, `vendor/`

## Dependencies

- `github.com/joho/godotenv` v1.5.1 — `.env` file loading

## Current Status

Phase 2 in progress — configuration module implemented (`internal/config/`). Remaining `internal/` packages are empty stubs.
