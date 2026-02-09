package main

import (
	"log/slog"
	"os"
	"github.com/DDRMin/GO-Backend/internal/env"
)


func main() {
	config := config{
		addr: ":8080",
		db: dbConfig{
			dbUrl: env.GetString("DB_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}

	api := API{
		config: config,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
	
}
