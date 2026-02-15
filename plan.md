# GoBlog Project Initialization — Implementation Plan

## Current State
- Repo has: `.gitignore` (default Go template), `LICENSE` (MIT), `.git/`
- Go version: `go1.25.6 linux/amd64`
- No `go.mod` exists yet

---

## Step 1: Initialize Go Module

Run `go mod init github.com/manasm11/goblog` to create `go.mod` with Go 1.25.6.

**Files created:**
- `go.mod` — will contain `module github.com/manasm11/goblog` and `go 1.25`

---

## Step 2: Create Directory Structure with `.gitkeep` Files

Create 18 directories, each with an empty `.gitkeep` file so Git tracks them:

| Directory | File |
|---|---|
| `cmd/server/` | *(will contain main.go — no .gitkeep needed)* |
| `internal/config/` | `internal/config/.gitkeep` |
| `internal/database/` | `internal/database/.gitkeep` |
| `internal/models/` | `internal/models/.gitkeep` |
| `internal/repository/` | `internal/repository/.gitkeep` |
| `internal/services/` | `internal/services/.gitkeep` |
| `internal/handlers/` | `internal/handlers/.gitkeep` |
| `internal/middleware/` | `internal/middleware/.gitkeep` |
| `internal/markdown/` | `internal/markdown/.gitkeep` |
| `internal/seo/` | `internal/seo/.gitkeep` |
| `templates/layouts/` | `templates/layouts/.gitkeep` |
| `templates/pages/` | `templates/pages/.gitkeep` |
| `templates/partials/` | `templates/partials/.gitkeep` |
| `templates/admin/` | `templates/admin/.gitkeep` |
| `static/css/` | `static/css/.gitkeep` |
| `static/js/` | `static/js/.gitkeep` |
| `static/images/` | `static/images/.gitkeep` |
| `uploads/` | `uploads/.gitkeep` |

**Note:** `cmd/server/` does NOT need a `.gitkeep` because it will contain `main.go`.

---

## Step 3: Create `cmd/server/main.go`

**File:** `cmd/server/main.go`

**Package:** `main`

**Imports:**
- `context`
- `encoding/json`
- `log`
- `net/http`
- `os`
- `os/signal`
- `syscall`
- `time`

**Structure:**

### `func main()`
1. Create a `http.NewServeMux()`
2. Register `GET /health` handler on the mux
3. Create `http.Server` with:
   - `Addr: ":8069"`
   - `Handler: mux`
4. Log `"GoBlog server starting on :8069"`
5. Start the server in a goroutine via `server.ListenAndServe()`
6. Create a channel listening for `os.Interrupt` and `syscall.SIGTERM` using `signal.NotifyContext` or `signal.Notify`
7. Block on the signal channel
8. On signal received, log shutdown message
9. Create a context with timeout (e.g., 10 seconds) for graceful shutdown
10. Call `server.Shutdown(ctx)`
11. Log completion and exit

### Health endpoint handler
- Respond to `GET /health` only (check `r.Method == http.MethodGet`; return 405 for other methods)
- Set `Content-Type: application/json`
- Write JSON: `{"status":"ok","version":"0.1.0"}`
- Use `json.NewEncoder(w).Encode()` or `json.Marshal` + `w.Write`

**Edge cases:**
- Non-GET requests to `/health` should return 405 Method Not Allowed
- Graceful shutdown should have a timeout so the process doesn't hang forever
- If `ListenAndServe` returns a non-`http.ErrServerClosed` error, log it fatally

---

## Step 4: Create `.env.example`

**File:** `.env.example`

**Contents (verbatim):**
```
GOBLOG_PORT=8069
GOBLOG_BASE_URL=http://localhost:8069
GOBLOG_BLOG_TITLE=Manas's Blog
GOBLOG_BLOG_DESCRIPTION=Thoughts on software engineering, technology, and life
GOBLOG_AUTHOR_NAME=Manas
GOBLOG_ADMIN_USERNAME=admin
GOBLOG_ADMIN_PASSWORD=changeme
GOBLOG_DB_PATH=./goblog.db
GOBLOG_UPLOAD_DIR=./uploads
GOBLOG_SESSION_SECRET=change-this-to-a-random-string
GOBLOG_ENV=development
```

---

## Step 5: Update `.gitignore`

**File:** `.gitignore` (replace existing contents)

The existing `.gitignore` has the default Go template. Replace it entirely with a project-specific version that includes the original useful entries plus the requested additions:

```gitignore
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out

# Build output
bin/
tmp/

# Dependencies
vendor/

# Database
goblog.db

# Environment
.env

# Uploads (keep directory via .gitkeep)
uploads/*
!uploads/.gitkeep

# Go workspace
go.work
go.work.sum
```

**Changes from existing:**
- Added `bin/`, `tmp/`
- Added `goblog.db`
- Added `uploads/*` and `!uploads/.gitkeep`
- Added `vendor/` (was commented out)
- Removed `coverage.*`, `*.coverprofile`, `profile.cov` (not in spec, but harmless — keep if desired)
- Kept essential Go ignores from original

---

## Step 6: Create `Makefile`

**File:** `Makefile`

**Targets:**

| Target | Command | Description |
|---|---|---|
| `run` | `go run cmd/server/main.go` | Run the server |
| `build` | `go build -o bin/goblog cmd/server/main.go` | Build binary |
| `test` | `go test ./... -v` | Run all tests |
| `templ` | `templ generate` | Generate templ templates |
| `dev` | `templ generate && go run cmd/server/main.go` | Dev workflow |
| `clean` | `rm -rf bin/ tmp/` | Clean build artifacts |

**Important details:**
- Each target should be declared `.PHONY` (none produce files matching their target name)
- Use tabs (not spaces) for recipe indentation — **critical for Makefile syntax**

---

## Step 7: Create `README.md`

**File:** `README.md`

**Sections:**

### Title & Description
- `# GoBlog`
- Brief description: A lightweight, SEO-friendly blog engine built with Go

### Tech Stack
- Go (HTTP server, backend logic)
- Templ (type-safe HTML templating)
- Pico CSS (minimal, classless CSS framework)
- SQLite (embedded database)

### Setup Instructions
1. Clone the repository
2. Copy `.env.example` to `.env` and customize values
3. Run `make dev` to start the development server
4. Visit `http://localhost:8069`

### Features (checkboxes — all unchecked)
- [ ] Markdown blog posts with frontmatter
- [ ] Admin dashboard for post management
- [ ] SEO-optimized HTML output (meta tags, Open Graph, sitemap)
- [ ] RSS/Atom feed generation
- [ ] Syntax highlighting for code blocks
- [ ] Image upload and management
- [ ] Tag/category support
- [ ] Responsive design with Pico CSS
- [ ] SQLite database with migrations
- [ ] Health check endpoint
- [ ] Graceful shutdown
- [ ] Session-based authentication

---

## Step 8: Run `go mod tidy`

Run `go mod tidy` after creating `main.go`. Since the code only uses standard library packages, this should be a no-op but ensures `go.sum` is clean and `go.mod` is valid.

---

## File Summary

| Action | File |
|---|---|
| **Create** | `go.mod` (via `go mod init`) |
| **Create** | `cmd/server/main.go` |
| **Create** | `.env.example` |
| **Modify** | `.gitignore` (replace contents) |
| **Create** | `Makefile` |
| **Create** | `README.md` |
| **Create** | 17 × `.gitkeep` files (all empty dirs except `cmd/server/`) |

**Total: 21 files created, 1 file modified**

---

## Execution Order

1. `go mod init github.com/manasm11/goblog`
2. Create all directories + `.gitkeep` files (parallel, no dependencies)
3. Create `cmd/server/main.go`
4. Create `.env.example`
5. Modify `.gitignore`
6. Create `Makefile`
7. Create `README.md`
8. Run `go mod tidy`

Steps 2–7 are independent and can be done in any order. Step 1 must come first (creates `go.mod`). Step 8 must come last (validates the module).
