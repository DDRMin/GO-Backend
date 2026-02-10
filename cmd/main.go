package main

import (
	"log/slog"
	"os"
	"github.com/DDRMin/GO-Backend/internal/env"
	"context"
	"github.com/jackc/pgx/v5"
)


func main() {
	ctx := context.Background()

	config := config{
		addr: ":8080",
		db: dbConfig{
			dbUrl: env.GetString("DB_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		},
	}
	
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn,err := pgx.Connect(ctx, config.db.dbUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Successfully connected to the database", "dbUrl", config.db.dbUrl)

	api := API{
		config: config,
	}


	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
	
}
