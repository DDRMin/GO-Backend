package main

import (
	"log/slog"
	"os"
)


func main() {
	config := config{
		addr: ":8080",
		db: dbConfig{
			dbUrl: "postgres://	user:password@localhost:5432/mydb",
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
