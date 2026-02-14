package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"semita.wk/go-pg-api/internal/config"
	"semita.wk/go-pg-api/internal/db"
	"semita.wk/go-pg-api/internal/repo"
	httpapi "semita.wk/go-pg-api/internal/http"
)
import "github.com/joho/godotenv"

func main() {
    _ = godotenv.Load() // reads .env into environment variables
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := db.NewPool(ctx, cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	wRepo := repo.NewWidgetRepo(pool)
	handlers := httpapi.NewHandlers(wRepo)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: httpapi.Routes(handlers),
	}

	go func() {
		log.Printf("listening on %s", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}
