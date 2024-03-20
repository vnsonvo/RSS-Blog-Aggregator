package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vnsonvo/RSS-Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't decode parameters: %s", err))
		return
	}

	newUser, err := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(newUser))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, req *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, req *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(req.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
