package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iamr-kumar/go-service/internal/databases"
)


func (config *apiConfig)handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newUser, err := config.DB.CreateUser(r.Context(), databases.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Email: params.Email,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not create user")
		return
	}

	respondWithJSON(w, http.StatusOK, convertDbUserToUser(newUser))
}