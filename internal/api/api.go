package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/api/controllers"
	"github.com/rm-ryou/sample_todo_app/internal/config"
	"github.com/rm-ryou/sample_todo_app/internal/repositories"
)

func Run(cfg *config.Config) {
	db, err := repositories.SetupConnection(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: controllers.InitRoutes(db),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
