package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mohammad-Mansoor/go-api/internal/config"
	"github.com/Mohammad-Mansoor/go-api/internal/handlers"
)

func main() {
	config.MustLoad()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Healthz)

	srv := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server Failed: %v", err)
	}

}
