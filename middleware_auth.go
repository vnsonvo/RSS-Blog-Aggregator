package main

import (
	"fmt"
	"net/http"

	"github.com/vnsonvo/RSS-Blog-Aggregator/internal/auth"
	"github.com/vnsonvo/RSS-Blog-Aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetAPIKey(req.Header)

		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, req, user)
	}
}
