# Implementation Plan: Configuration Module (`internal/config`)

## Overview

Create the configuration module at `internal/config/config.go` that loads configuration from environment variables (with `.env` file support via `godotenv`), applies defaults, and validates required fields. Add comprehensive tests in `internal/config/config_test.go`. Run `go mod tidy` to fetch the new dependency.

---

## Files to Create/Modify

### 1. Create `internal/config/config.go`

**Package:** `config`

**Imports:** `log`, `os`, `github.com/joho/godotenv`

#### Struct: `Config`

```go
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
```

All fields are exported strings. No struct tags needed (this is not deserialized from JSON/YAML).

#### Function: `Load() *Config`

Signature: `func Load() *Config`

Implementation steps:

1. Call `godotenv.Load()` — **ignore the returned error**. The `.env` file is optional. If it exists, `godotenv` loads its key-value pairs into the process environment. Crucially, `godotenv.Load()` does **not** overwrite existing env vars, so real environment variables always take precedence over `.env` values.

2. Construct the `Config` struct using the `getEnv` helper for each field:

   | Field             | Env Var                    | Default                    |
   |-------------------|----------------------------|----------------------------|
   | `Port`            | `GOBLOG_PORT`              | `"8069"`                   |
   | `BaseURL`         | `GOBLOG_BASE_URL`          | `"http://localhost:8069"`  |
   | `BlogTitle`       | `GOBLOG_BLOG_TITLE`        | `"Manas's Blog"`           |
   | `BlogDescription` | `GOBLOG_BLOG_DESCRIPTION`  | `""` (empty)               |
   | `AuthorName`      | `GOBLOG_AUTHOR_NAME`       | `"Manas"`                  |
   | `AdminUsername`    | `GOBLOG_ADMIN_USERNAME`    | `""` (empty)               |
   | `AdminPassword`   | `GOBLOG_ADMIN_PASSWORD`    | `""` (empty)               |
   | `DBPath`          | `GOBLOG_DB_PATH`           | `"./goblog.db"`            |
   | `UploadDir`       | `GOBLOG_UPLOAD_DIR`        | `"./uploads"`              |
   | `SessionSecret`   | `GOBLOG_SESSION_SECRET`    | `""` (empty)               |
   | `Env`             | `GOBLOG_ENV`               | `"development"`            |

3. **Validate `Env`:** If `cfg.Env` is neither `"development"` nor `"production"`, call `log.Fatalf("invalid GOBLOG_ENV value %q: must be \"development\" or \"production\"", cfg.Env)`.

4. **Validate admin credentials in production:** If `cfg.Env == "production"` and (`cfg.AdminUsername == ""` or `cfg.AdminPassword == ""`), call `log.Fatalf("GOBLOG_ADMIN_USERNAME and GOBLOG_ADMIN_PASSWORD are required in production")`.

5. Return `&cfg`.

#### Helper: `getEnv(key, fallback string) string`

Signature: `func getEnv(key, fallback string) string`

- Unexported (lowercase).
- Calls `os.Getenv(key)`.
- If the result is `""`, return `fallback`.
- Otherwise return the value from the environment.

**Note:** `os.Getenv` returns `""` for both unset and explicitly-empty env vars. This means setting `GOBLOG_PORT=""` would cause the default `"8069"` to be used. This is acceptable and standard Go behavior.

---

### 2. Create `internal/config/config_test.go`

**Package:** `config` (same package for access to unexported `getEnv`)

**Imports:** `os`, `testing`

#### Helper: clearing all GOBLOG env vars

Create a test helper (unexported) that unsets all known `GOBLOG_*` variables to ensure test isolation:

```
func clearEnv(t *testing.T)
```

Calls `t.Setenv(key, "")` would set them to empty (which `getEnv` treats as unset), but to be fully clean, use `os.Unsetenv` for each key inside a `t.Cleanup`. The list of keys to unset:
- `GOBLOG_PORT`
- `GOBLOG_BASE_URL`
- `GOBLOG_BLOG_TITLE`
- `GOBLOG_BLOG_DESCRIPTION`
- `GOBLOG_AUTHOR_NAME`
- `GOBLOG_ADMIN_USERNAME`
- `GOBLOG_ADMIN_PASSWORD`
- `GOBLOG_DB_PATH`
- `GOBLOG_UPLOAD_DIR`
- `GOBLOG_SESSION_SECRET`
- `GOBLOG_ENV`

**Better approach:** Use `t.Setenv` for the keys we want to set in each test. `t.Setenv` automatically saves and restores the original value after the test, providing isolation without manual cleanup.

#### Test: `TestLoadDefaults(t *testing.T)`

- Unset all `GOBLOG_*` env vars (loop through each key, call `os.Unsetenv`).
- Call `cfg := Load()`.
- Assert each field matches its default value using direct comparison with `t.Errorf`:
  - `cfg.Port == "8069"`
  - `cfg.BaseURL == "http://localhost:8069"`
  - `cfg.BlogTitle == "Manas's Blog"`
  - `cfg.BlogDescription == ""`
  - `cfg.AuthorName == "Manas"`
  - `cfg.AdminUsername == ""`
  - `cfg.AdminPassword == ""`
  - `cfg.DBPath == "./goblog.db"`
  - `cfg.UploadDir == "./uploads"`
  - `cfg.SessionSecret == ""`
  - `cfg.Env == "development"`

**Note:** This test implicitly validates that `Load()` does NOT `log.Fatal` in development mode without admin credentials.

#### Test: `TestLoadFromEnvVars(t *testing.T)`

- Use `t.Setenv` for every `GOBLOG_*` variable with non-default values:
  - `GOBLOG_PORT` = `"9090"`
  - `GOBLOG_BASE_URL` = `"https://example.com"`
  - `GOBLOG_BLOG_TITLE` = `"Test Blog"`
  - `GOBLOG_BLOG_DESCRIPTION` = `"A test blog"`
  - `GOBLOG_AUTHOR_NAME` = `"Tester"`
  - `GOBLOG_ADMIN_USERNAME` = `"admin"`
  - `GOBLOG_ADMIN_PASSWORD` = `"secret"`
  - `GOBLOG_DB_PATH` = `"/tmp/test.db"`
  - `GOBLOG_UPLOAD_DIR` = `"/tmp/uploads"`
  - `GOBLOG_SESSION_SECRET` = `"mysecret"`
  - `GOBLOG_ENV` = `"production"`
- Call `cfg := Load()`.
- Assert each field matches the env var value, not the default.

#### Test: `TestLoadPartialEnvVars(t *testing.T)`

- Use `t.Setenv` for only a subset:
  - `GOBLOG_PORT` = `"3000"`
  - `GOBLOG_BLOG_TITLE` = `"Custom Blog"`
- Unset the rest (to be safe, unset all other `GOBLOG_*` vars).
- Call `cfg := Load()`.
- Assert:
  - `cfg.Port == "3000"` (overridden)
  - `cfg.BlogTitle == "Custom Blog"` (overridden)
  - `cfg.BaseURL == "http://localhost:8069"` (default)
  - `cfg.Env == "development"` (default)
  - All other fields at defaults.

#### Test: `TestGetEnv(t *testing.T)`

- **Subtest "returns env var value when set":**
  - `t.Setenv("TEST_CONFIG_KEY", "myvalue")`
  - Assert `getEnv("TEST_CONFIG_KEY", "default") == "myvalue"`

- **Subtest "returns fallback when not set":**
  - `os.Unsetenv("TEST_CONFIG_KEY_MISSING")`
  - Assert `getEnv("TEST_CONFIG_KEY_MISSING", "fallback") == "fallback"`

- **Subtest "returns fallback when set to empty":**
  - `t.Setenv("TEST_CONFIG_KEY_EMPTY", "")`
  - Assert `getEnv("TEST_CONFIG_KEY_EMPTY", "default") == "default"`

#### Validation Tests (not included — rationale)

The validation logic (`log.Fatalf` for invalid `Env` or missing admin credentials in production) calls `os.Exit(1)` under the hood. Testing this requires either:
- Running a subprocess via `exec.Command` and checking the exit code
- Refactoring to inject a logger or return an error

Neither is requested. The existing tests implicitly prove the happy paths work (development mode without admin creds succeeds, production mode with admin creds succeeds). A comment in the test file should note that fatal validation paths are exercised by the design of the other tests but not explicitly asserted.

---

### 3. `go.mod` and `go.sum` (modified by `go mod tidy`)

Running `go mod tidy` will:
- Add `require github.com/joho/godotenv v1.5.1` (or latest) to `go.mod`
- Create `go.sum` with the checksum for the dependency

No manual edits to `go.mod` are needed.

---

## Files NOT Modified

- **`cmd/server/main.go`** — Not in scope. A future step will integrate `config.Load()` into `main()` to replace the hardcoded port.
- **`Makefile`** — No changes needed.
- **`.env.example`** — Already contains all the correct `GOBLOG_*` variable names that map to the `Config` struct fields.

---

## Edge Cases

| Edge Case | Behavior |
|---|---|
| `.env` file does not exist | `godotenv.Load()` error is silently ignored; falls back to real env vars + defaults |
| `.env` file is malformed | Same as above — error ignored, real env vars + defaults used |
| `GOBLOG_ENV` set to invalid value (e.g., `"staging"`) | `log.Fatalf` with descriptive message |
| Production mode, missing `AdminUsername` | `log.Fatalf` requiring both admin credentials |
| Production mode, missing `AdminPassword` | `log.Fatalf` requiring both admin credentials |
| Development mode, missing admin credentials | Allowed — `Load()` succeeds |
| Env var set to empty string (`GOBLOG_PORT=""`) | Treated as unset; default value used (standard Go `os.Getenv` behavior) |
| Real env var + `.env` both set same key | Real env var wins (`godotenv.Load` does not overwrite existing vars) |

---

## Execution Order

1. Create `internal/config/config.go`
2. Create `internal/config/config_test.go`
3. Run `go mod tidy` (fetches `godotenv`, creates `go.sum`)
4. Run `make test` to verify all tests pass
