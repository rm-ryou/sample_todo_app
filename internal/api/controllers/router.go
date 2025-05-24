package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rm-ryou/sample_todo_app/internal/api/controllers/presenter/response"
	"github.com/rm-ryou/sample_todo_app/internal/repositories"
	"github.com/rm-ryou/sample_todo_app/internal/services"
	"github.com/rs/cors"
)

// FIXME: Avoid initializing service, repository, controller in InitRouter
// TODO: Use middleware and frameworks such as gin and echo for easy routing configuration
func InitRoutes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/health", healthCheckMux())
	mux.Handle("/v1/rooms/", roomMux(db))
	mux.Handle("/v1/boards/{boardId}/todos/", todoMux(db))

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

func roomMux(db *sql.DB) *http.ServeMux {
	repo := repositories.NewRoomRepository(db)
	service := services.NewRoomService(repo)
	controller := NewRoomController(service)

	mux := http.NewServeMux()
	mux.Handle("/v1/rooms/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetAll(w, r)
		case http.MethodPost:
			controller.Create(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, fmt.Errorf("%s is not allowed", r.Method))
		}
	}))
	mux.Handle("/v1/rooms/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controller.Update(w, r)
		case http.MethodDelete:
			controller.Delete(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, fmt.Errorf("%s is not allowed", r.Method))
		}
	}))

	return mux
}

func todoMux(db *sql.DB) *http.ServeMux {
	repository := repositories.NewTodoRepository(db)
	service := services.NewTodoService(repository)
	controller := NewTodoController(service)

	mux := http.NewServeMux()
	mux.Handle("/v1/boards/{boardId}/todos/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.Create(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, fmt.Errorf("%s is not allowed", r.Method))
		}
	}))
	mux.Handle("/v1/boards/{boardId}/todos/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetById(w, r)
		case http.MethodPut:
			controller.Update(w, r)
		case http.MethodDelete:
			controller.Delete(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, fmt.Errorf("%s is not allowed", r.Method))
		}
	}))

	return mux
}
