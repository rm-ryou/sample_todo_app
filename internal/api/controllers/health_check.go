package controllers

import (
	"net/http"

	"github.com/rm-ryou/sample_todo_app/internal/api/controllers/presenter/response"
)

type HealthCheck struct{}

func (HealthCheck) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	response.Basic(w, http.StatusOK, response.BasicResponse{Message: "ok"})
}
