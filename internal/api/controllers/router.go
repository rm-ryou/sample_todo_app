package controllers

import (
	"database/sql"
	"net/http"

	"github.com/rm-ryou/sample_todo_app/internal/repositories"
	"github.com/rm-ryou/sample_todo_app/internal/services"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/health", healthCheckMux())
	mux.Handle("/v1/todos/", todoMux(db))

	return mux
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
	mux.Handle("GET /v1/todos/{id}", http.HandlerFunc(controller.GetById))
	mux.Handle("POST /v1/todos/", http.HandlerFunc(controller.Create))
	mux.Handle("PUT /v1/todos/{id}", http.HandlerFunc(controller.Update))
	mux.Handle("DELETE /v1/todos/{id}", http.HandlerFunc(controller.Delete))

	return mux
}
