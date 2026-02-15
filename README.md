# GoBlog

A personal blogging platform built with Go, Templ, Pico CSS, and SQLite.

## Tech Stack

- **Go** — backend server and routing
- **Templ** — type-safe HTML templating
- **Pico CSS** — minimal, classless CSS framework
- **SQLite** — embedded database

## Project Structure

```
cmd/server/         — application entry point
internal/
  config/           — configuration loading
  database/         — SQLite database setup and migrations
  models/           — data models
  repository/       — data access layer
  services/         — business logic
  handlers/         — HTTP route handlers
  middleware/       — HTTP middleware
  markdown/         — markdown rendering
  seo/              — sitemap, RSS, structured data
templates/          — Templ templates (layouts, pages, partials, admin)
static/             — CSS, JS, images
uploads/            — user-uploaded files
```

## Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
2. Edit `.env` with your desired settings.
3. Run the development server:
   ```bash
   make dev
   ```

## Make Targets

| Target  | Description                              |
|---------|------------------------------------------|
| `run`   | `go run cmd/server/main.go`              |
| `build` | Build binary to `bin/goblog`             |
| `test`  | Run all tests with verbose output        |
| `templ` | Generate Templ templates                 |
| `dev`   | Generate templates and run server        |
| `clean` | Remove `bin/` and `tmp/` directories     |

## Environment Variables

See `.env.example` for all available variables. Key settings:

| Variable              | Default                  | Description              |
|-----------------------|--------------------------|--------------------------|
| `GOBLOG_PORT`         | `8069`                   | Server port              |
| `GOBLOG_DB_PATH`      | `./goblog.db`            | SQLite database path     |
| `GOBLOG_ENV`          | `development`            | Environment mode         |
| `GOBLOG_UPLOAD_DIR`   | `./uploads`              | Upload directory         |

## API

### Health Check

```
GET /health
```

Returns:
```json
{"status": "ok", "version": "0.1.0"}
```

## Features

- [ ] Markdown blog posts with frontmatter
- [ ] Admin dashboard with authentication
- [ ] Post CRUD (create, read, update, delete)
- [ ] Tag and category support
- [ ] RSS feed generation
- [ ] SEO meta tags and Open Graph support
- [ ] Sitemap generation
- [ ] Image uploads
- [ ] Syntax highlighting for code blocks
- [ ] Responsive design with Pico CSS
- [ ] Draft/published post states
- [ ] Full-text search
