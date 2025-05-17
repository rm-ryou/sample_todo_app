package controllers

import (
	"net/http"
)

type HealthCheck struct{}

func (HealthCheck) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	CommonResponse(w, http.StatusOK, "ok")
}
