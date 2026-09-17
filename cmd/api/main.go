package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Mohammad-Mansoor/go-api/internal/config"
	"github.com/Mohammad-Mansoor/go-api/internal/db"
	"github.com/Mohammad-Mansoor/go-api/internal/handlers"
	"github.com/Mohammad-Mansoor/go-api/internal/middlewares"
)

func main() {
	ctx := config.MustLoad()
	db, err := db.ConnectDB(ctx.DB_URL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	lh := handlers.NewListingHandlers(db)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Healthz)
	mux.HandleFunc("GET /listings", lh.Listing)
	mux.HandleFunc("DELETE /listing/{id}", lh.DeleteListing)
	handler := middlewares.RequestId(mux)
	srv := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server Failed: %v", err)
	}

}
