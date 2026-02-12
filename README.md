# GO-Backend Service

Minimal Go HTTP API starter with Chi, PostgreSQL, Atlas migrations, and sqlc code generation.

## Overview

This project is a lightweight API server built with Go 1.25.5, designed to be simple yet robust. Key features include:

-   **Router**: `chi` v5 for fast and composable routing.
-   **Database**: `pgx` v5 for high-performance PostgreSQL driver and connection pooling.
-   **Migrations**: Managed by [Atlas](https://atlasgo.io) for declarative schema management (`internal/adapters/migrations/schema.hcl`).
-   **Code Generation**: [sqlc](https://sqlc.dev) for type-safe SQL queries (`internal/adapters/sqlc`).
-   **Logging**: Structured logging using Go's standard library `slog`.
-   **Task Runner**: `Taskfile.yml` for simplified command execution.

## Project Structure

```bash
.
├── cmd/                # Application entrypoints
│   ├── main.go         # Service entrypoint
│   └── api.go          # HTTP server and router setup
├── internal/
│   ├── adapters/       # Interface adapters (database, migrations)
│   │   ├── migrations/ # Atlas schema definition (schema.hcl) and SQL dump (schema.sql)
│   │   └── sqlc/       # sqlc generated Go code and SQL queries
│   ├── env/            # Environment variable utilities
│   ├── json/           # JSON response helpers
│   └── products/       # Domain logic for product management
├── atlas.hcl           # Atlas configuration file
├── sqlc.yml            # sqlc configuration file
└── Taskfile.yml        # Task runner definitions
```

## Prerequisites

Ensure you have the following installed:

-   **Go**: 1.25.5 or higher
-   **PostgreSQL**: 16 or higher (or Docker)
-   **Atlas CLI**: For managing database schemas
    -   MacOS: `brew install ariga/tap/atlas`
    -   Windows/Linux: Download from [ariga.io](https://atlasgo.io/getting-started)
-   **Task**: (Optional) For running `Taskfile.yml` commands. Install via [taskfile.dev](https://taskfile.dev/installation/).
-   **sqlc**: (Optional) For regenerating database code. Install via [sqlc.dev](https://docs.sqlc.dev/en/latest/overview/install.html).

## Getting Started

1.  **Clone the repository**

    ```bash
    git clone https://github.com/DDRMin/GO-Backend.git
    cd GO-Backend
    ```

2.  **Environment Setup**

    Configure your database connection string. You can set the `DB_URL` environment variable or use a `.env` file if you implement one.

    **Windows (PowerShell):**
    ```powershell
    $env:DB_URL="postgres://user:password@localhost:5432/mydb?sslmode=disable"
    ```

    **Mac/Linux:**
    ```bash
    export DB_URL="postgres://user:password@localhost:5432/mydb?sslmode=disable"
    ```

3.  **Install Dependencies**

    ```bash
    go mod tidy
    ```

4.  **Database Setup (Atlas)**

    Apply the database schema defined in `internal/adapters/migrations/schema.hcl`.

    ```bash
    # Using Taskfile (Recommended)
    task atlas-apply

    # Manual command
    atlas schema apply --env local
    ```

    This command reads `DB_URL`, compares the declarative schema with the database state, and applies necessary changes.

5.  **Run the Application**

    ```bash
    go run ./cmd
    ```

    The server will start on port `8080` (default).

## Development Workflow

### Database Schema Changes

1.  Modify `internal/adapters/migrations/schema.hcl`.
2.  Apply changes to your local database:
    ```bash
    task atlas-apply
    ```
    This also updates `internal/adapters/migrations/schema.sql` which serves as the source of truth for `sqlc`.

### Generating SQL Code

If you add or modify SQL queries in `internal/adapters/sqlc/queries.sql`:

1.  Run the generator:
    ```bash
    task sqlc
    # OR
    sqlc generate
    ```
2.  Use the generated methods in your Go code (see `internal/products/service.go` for examples).

## API Endpoints

| Method | Endpoint    | Description                                      |
| :----- | :---------- | :----------------------------------------------- |
| `GET`  | `/health`   | Health check. Returns 200 OK or 503 Unavailable. |
| `GET`  | `/products` | Retreives a list of products.                    |

## Configuration

The application is configured via environment variables:

| Variable | Description                                             | Default     |
| :------- | :------------------------------------------------------ | :---------- |
| `DB_URL` | PostgreSQL connection string                            | (Required)  |
| `PORT`   | Port for the HTTP server (if configured in `config.go`) | `:8080`     |
