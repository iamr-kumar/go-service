package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iamr-kumar/go-service/internal/databases"
)


func (config *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user databases.User) {
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := config.DB.CreateFeed(r.Context(), databases.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not create user") 
		return
	}

	respondWithJSON(w, http.StatusCreated, convertDbFeedToFeed(feed))
}

func (config *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := config.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, convertDbFeedsToFeed(feeds))
}