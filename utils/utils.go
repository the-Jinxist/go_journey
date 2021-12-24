package utils

import (
	"encoding/json"
	"go_journey/models"
	"net/http"
)

func SendError(w http.ResponseWriter, error models.Error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func SendSuccess(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
