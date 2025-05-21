package controllers

import (
	"database/sql"
	"net/http"

	"github.com/rm-ryou/sample_todo_app/internal/repositories"
	"github.com/rm-ryou/sample_todo_app/internal/services"
	"github.com/rs/cors"
)

func InitRoutes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/health", healthCheckMux())
	mux.Handle("/v1/todos/", todoMux(db))

	c := cors.New(cors.Options{
		// TODO: fix allow origin
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	return c.Handler(mux)
}

func healthCheckMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/health", HealthCheck{})

	return mux
}

func todoMux(db *sql.DB) *http.ServeMux {
	repository := repositories.NewTodoRepository(db)
	service := services.NewTodoService(repository)
	controller := NewTodoController(service)

	mux := http.NewServeMux()
	mux.Handle("GET /v1/todos/", http.HandlerFunc(controller.GetAll))
	mux.Handle("GET /v1/todos/{id}", http.HandlerFunc(controller.GetById))
	mux.Handle("POST /v1/todos/", http.HandlerFunc(controller.Create))
	mux.Handle("PUT /v1/todos/{id}", http.HandlerFunc(controller.Update))
	mux.Handle("DELETE /v1/todos/{id}", http.HandlerFunc(controller.Delete))

	return mux
}
