package router

import (
	"database/sql"
	"net/http"

	"github.com/rm-ryou/sample_todo_app/internal/api/handler"
	"github.com/rm-ryou/sample_todo_app/internal/repository"
	"github.com/rm-ryou/sample_todo_app/internal/service/todo"
)

func New(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/health", healthCheckMux())
	mux.Handle("/v1/todos/", todoMux(db))

	return mux
}

func healthCheckMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/health", handler.HealthCheck{})

	return mux
}

func todoMux(db *sql.DB) *http.ServeMux {
	repository := repository.NewTodo(db)
	service := todo.NewService(repository)
	todoHandler := handler.NewTodo(service)

	mux := http.NewServeMux()
	mux.Handle("GET /v1/todos/{id}", http.HandlerFunc(todoHandler.GetById))
	mux.Handle("POST /v1/todos/", http.HandlerFunc(todoHandler.CreateTodo))
	mux.Handle("PUT /v1/todos/{id}", http.HandlerFunc(todoHandler.Update))
	mux.Handle("DELETE /v1/todos/{id}", http.HandlerFunc(todoHandler.Delete))

	return mux
}
