package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iamr-kumar/go-service/internal/databases"
)


func (config *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user databases.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedFollows, err := config.DB.CreateFeedFollow(r.Context(), databases.CreateFeedFollowParams{
		ID: uuid.New(),
		FeedID: params.FeedID,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not create feed follow") 
		return 
	}

	respondWithJSON(w, http.StatusCreated, convertDbFeedFollowToFeedFollow(feedFollows))
} 

func(config *apiConfig) getFeedFollows(w http.ResponseWriter, r *http.Request, user databases.User) {
	feedFollows, err := config.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, convertDbFeedFollowsToFeedFollows(feedFollows))
}