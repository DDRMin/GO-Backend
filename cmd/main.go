package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DDRMin/GO-Backend/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config := config{
		addr: ":8080",
		db: dbConfig{
			dbUrl: env.GetString("DB_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}

	pool, err := pgxpool.New(ctx, config.db.dbUrl)
	if err != nil {
		slog.Error("Failed to create connection pool", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	slog.Info("Connected to database", "url", config.db.dbUrl)

	api := API{
		config: config,
		pool:   pool,
	}

	if err := api.run(ctx, api.mount()); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}

	slog.Info("Application shut down gracefully")
}
