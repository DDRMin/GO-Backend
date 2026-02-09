package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DDRMin/GO-Backend/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	config config
}

func (app *API) mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID) 
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	productsHandler := products.NewHandler(nil)
	router.Get("/products", productsHandler.ListProducts)

	return router
}

func (app *API) run(h http.Handler) error {
	server := &http.Server{
		Addr:    app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return server.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbUrl string
}