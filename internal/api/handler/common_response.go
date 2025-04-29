package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func CommonResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	res := response{Message: msg}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	log.Printf("Error: %v", err)
	CommonResponse(w, code, http.StatusText(code))
}
