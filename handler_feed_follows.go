package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vnsonvo/RSS-Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't decode parameters: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follows: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollows(req.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feed follows: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	feedFollowIDStr := req.PathValue("feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow id: %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follows: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
