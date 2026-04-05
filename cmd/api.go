package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Piyush-Singh-coder/ecom/internal/adapters/postgresql/sqlc"
	"github.com/Piyush-Singh-coder/ecom/internal/orders"
	"github.com/Piyush-Singh-coder/ecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	//logger
	db *pgx.Conn
}

// mount

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60*time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	
	newService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(newService)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/products/{id}", productHandler.ProductById)

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)
	r.Post("/orders", orderHandler.PlaceOrder)

	return r
}

// run

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr : app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
