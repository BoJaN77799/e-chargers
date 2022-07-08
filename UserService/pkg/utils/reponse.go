package utils

import (
	"encoding/json"
	"net/http"
)

func BadRequestResponse(w http.ResponseWriter, err string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func CreatedResponse(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func OKResponse(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
