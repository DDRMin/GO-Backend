<div align="center">

# GO-Backend Service

Minimal Go HTTP API starter with Chi, PostgreSQL, and sqlc code generation.

</div>

## Overview

- Lightweight API server built with Chi, `pgx`, and Go 1.25.5.
- Connects to PostgreSQL via `DB_URL` (defaults to a local connection).
- Product domain stub with request logging, timeouts, and structured logging via `slog`.
- sqlc-backed queries for `products`; migrations stored under `internal/adapters/migrations`.

## Project Layout

- `cmd/` – application entrypoint (`main.go`) and router wiring (`api.go`).
- `internal/products/` – product handler and service interface (currently returns an empty list).
- `internal/adapters/sqlc/` – sqlc queries and generated code for `products`.
- `internal/adapters/migrations/` – database migrations (initial products table).
- `internal/env/` – environment variable helpers.
- `internal/json/` – JSON response helper.

## Quick Start

1) **Prerequisites**

- Go 1.25.5+
- PostgreSQL
- Optional: `migrate` CLI (for running migrations) and `sqlc` (for regenerating query code)

2) **Environment**

Set the database URL or rely on the default:

```bash
set DB_URL=postgres://user:password@localhost:5432/mydb?sslmode=disable   # PowerShell: $env:DB_URL="..."
```

3) **Install dependencies**

```bash
go mod tidy
```

4) **Run the API**

```bash
go run ./cmd
```

The server listens on `:8080`.

## API

- `GET /products` – returns an empty `products` array placeholder.
- `GET /` – health/root check.

## Database

Migrations live in `internal/adapters/migrations` (initial products table). Run with `migrate` CLI:

```bash
migrate -path internal/adapters/migrations -database "%DB_URL%" up
migrate -path internal/adapters/migrations -database "%DB_URL%" down
```

## sqlc

Generate query interfaces and models from `internal/adapters/sqlc/queries.sql`:

```bash
sqlc generate
```

## Local Development Tips

- Logging: structured logs via `slog` print to stdout.
- Timeouts: Chi middleware sets a 60s request timeout; HTTP server uses 30s read/write timeouts.
- Testing: run `go test ./...` once tests are added.
