package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
	"github.com/DDRMin/GO-Backend/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

const shutdownTimeout = 30 * time.Second

type API struct {
	config config
	pool   *pgxpool.Pool
}

func (app *API) mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/health", app.healthCheck)

	productService := products.NewService(repo.New(app.pool))
	productsHandler := products.NewHandler(productService)
	router.Get("/products", productsHandler.ListProducts)
	router.Get("/products/{id}", productsHandler.GetProduct)
	router.Post("/products", productsHandler.CreateProduct)

	return router
}

func (app *API) run(ctx context.Context, h http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	slog.Info("Server starting", "addr", app.config.addr)

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		slog.Info("Shutdown signal received, draining connections...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			slog.Error("Forced shutdown", "error", err)
			return err
		}

		slog.Info("Server stopped gracefully")
		return nil

	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (app *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	if err := app.pool.Ping(r.Context()); err != nil {
		slog.Error("Health check failed", "error", err)
		http.Error(w, "unhealthy", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbUrl string
}
