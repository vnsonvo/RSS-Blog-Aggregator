package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	payload := struct {
		Status string `json:"status"`
	}{Status: "ok"}

	respondWithJSON(w, http.StatusOK, payload)

}

func handlerErr(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, http.StatusInternalServerError,
		"Internal Server Error")
}
