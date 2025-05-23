package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type BasicResponse struct {
	Message string `json:"message"`
}

func Basic(w http.ResponseWriter, code int, resItem any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(resItem); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, code int, err error) {
	log.Printf("Error output by controller: %v", err)
	res := BasicResponse{Message: http.StatusText(code)}

	Basic(w, code, res)
}
