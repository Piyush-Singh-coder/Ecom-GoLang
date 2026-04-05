package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Piyush-Singh-coder/ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user=postgres password=postgres host=localhost port=5433 dbname=ecom sslmode=disable"),
		},
	}


	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)


	// Database connection
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	
	logger.Info("database connection successful", "dsn", cfg.db.dsn)


	api := application{
		config: cfg,
		db: conn,
	}


	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}