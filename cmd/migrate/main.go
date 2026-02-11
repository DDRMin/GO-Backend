package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/DDRMin/GO-Backend/internal/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultTimeout = 10 * time.Second
)

func main() {
	var (
		cmd   = flag.String("cmd", "", "up / down / version")
		steps = flag.Int("steps", 0, "number of steps to migrate (optional for up/down)")
	)
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// In a real app, you might want to load this from a .env file or config service
	dbURL := env.GetString("DB_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable")
	sourceURL := "file://internal/adapters/migrations"

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		logger.Error("Failed to create migrate instance", "error", err)
		os.Exit(1)
	}
	defer m.Close()

	initialVersion, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		logger.Error("Failed to get current migration version", "error", err)
		os.Exit(1)
	}

	if dirty {
		logger.Warn("Database is dirty. Forcing version to clean state before proceeding...", "version", initialVersion)
		// If dirty, we might want to force the version to the previous one to retry, or letting user handle it manually.
		// For safety in this CLI, we will just warn. Often 'm.Force(version)' is used to fix dirty state manually.
		logger.Warn("Please fix the dirty state manually using the 'force' command in 'migrate' CLI or by database inspection.")
	}

	logger.Info("Current migration version", "version", initialVersion, "dirty", dirty)

	start := time.Now()

	switch *cmd {
	case "up":
		if *steps > 0 {
			err = m.Steps(*steps)
		} else {
			err = m.Up()
		}
	case "down":
		if *steps > 0 {
			err = m.Steps(-*steps)
		} else {
			err = m.Down()
		}
	case "version":
		// Already logged above
		return
	default:
		logger.Error("Unknown command", "cmd", *cmd)
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("No changes to apply")
			return
		}
		logger.Error("Migration failed", "error", err, "duration", time.Since(start))
		os.Exit(1)
	}

	finalVersion, _, _ := m.Version()
	logger.Info("Migration finished successfully", "cmd", *cmd, "from", initialVersion, "to", finalVersion, "duration", time.Since(start))
}
