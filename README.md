<div align="center">

# GO-Backend Service

Minimal Go HTTP API starter with Chi, PostgreSQL, Atlas migrations, and sqlc code generation.

</div>

## Overview

- Lightweight API server built with Chi, `pgx`, and Go 1.25.5.
- Connects to PostgreSQL via `DB_URL` (defaults to a local connection) and pings the DB on `/health`.
- Database schema managed with Atlas from `internal/adapters/migrations/schema.hcl`; `schema.sql` stays in sync for sqlc.
- Product domain stub with request logging, timeouts, and structured logging via `slog`.
- sqlc-backed queries for `products` (currently unused in the handler).

## Project Layout

- `cmd/` – application entrypoint (`main.go`) and router wiring (`api.go`).
- `internal/products/` – product handler and service (returns an empty list placeholder).
- `internal/adapters/sqlc/` – sqlc queries and generated code for `products`.
- `internal/adapters/migrations/` – Atlas schema (`schema.hcl`) and companion `schema.sql` for sqlc alignment.
- `internal/env/` – environment variable helpers.
- `internal/json/` – JSON response helper.

## Quick Start

1) **Prerequisites**
	- Go 1.25.5+
	- PostgreSQL
	- Atlas CLI (`brew install ariga/tap/atlas` or download from ariga.io)
	- Optional: `sqlc` for regenerating query code

2) **Environment**
	Set the database URL or rely on the default:

	```bash
	set DB_URL=postgres://user:password@localhost:5432/mydb?sslmode=disable   # PowerShell: $env:DB_URL="..."
	```

3) **Install dependencies**
	```bash
	go mod tidy
	```

4) **Apply the schema with Atlas**
	```bash
	atlas schema apply --env local
	```
	The `local` env in `atlas.hcl` reads `DB_URL`, formats and plans against the dev container image (`postgres:17`), then applies the plan to your database.

5) **Run the API**
	```bash
	go run ./cmd
	```
	The server listens on `:8080`.

## API

- `GET /health` – pings the database connection pool; returns `503` if unreachable.
- `GET /products` – returns an empty `products` array placeholder.

## Database (Atlas)

- Schema source of truth: `internal/adapters/migrations/schema.hcl` (referenced by `atlas.hcl`).
- Apply changes: `atlas schema apply --env local`.
- Inspect current DB: `atlas schema inspect -u "$env:DB_URL"` (PowerShell) or `atlas schema inspect -u "$DB_URL"` (bash).
- Keep `schema.sql` aligned with `schema.hcl` for sqlc (see comment inside the file).

## sqlc

Generate query interfaces and models from `internal/adapters/sqlc/queries.sql`:

```bash
sqlc generate
```

## Local Development Tips

- Logging: structured logs via `slog` print to stdout.
- Timeouts: Chi middleware sets a 60s request timeout; HTTP server uses 30s read/write timeouts.
- Testing: run `go test ./...` once tests are added.
