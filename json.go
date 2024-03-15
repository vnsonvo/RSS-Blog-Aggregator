package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Printf("Responding with 5xx error: %s", msg)
	}

	respondWithJSON(w, statusCode, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}
