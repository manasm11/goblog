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

## Configuration

Configuration is loaded in order of priority: `.env` file (if present) → environment variables → defaults. See `.env.example` for a template.

| Variable                   | Default                  | Description                                      |
|----------------------------|--------------------------|--------------------------------------------------|
| `GOBLOG_PORT`              | `8069`                   | Server port                                      |
| `GOBLOG_BASE_URL`          | `http://localhost:8069`  | Public base URL                                  |
| `GOBLOG_BLOG_TITLE`        | `Manas's Blog`           | Blog title                                       |
| `GOBLOG_BLOG_DESCRIPTION`  | *(empty)*                | Blog description                                 |
| `GOBLOG_AUTHOR_NAME`       | `Manas`                  | Author name                                      |
| `GOBLOG_ADMIN_USERNAME`    | *(empty)*                | Admin login username (required in production)     |
| `GOBLOG_ADMIN_PASSWORD`    | *(empty)*                | Admin login password (required in production)     |
| `GOBLOG_DB_PATH`           | `./goblog.db`            | SQLite database path                             |
| `GOBLOG_UPLOAD_DIR`        | `./uploads`              | Upload directory                                 |
| `GOBLOG_SESSION_SECRET`    | *(empty)*                | Session signing secret                           |
| `GOBLOG_ENV`               | `development`            | `development` or `production`                    |

**Note:** In production mode, `GOBLOG_ADMIN_USERNAME` and `GOBLOG_ADMIN_PASSWORD` must be set or the server will refuse to start.

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
