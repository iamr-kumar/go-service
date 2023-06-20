package main

import (
	"net/http"

	"github.com/iamr-kumar/go-service/internal/auth"
	"github.com/iamr-kumar/go-service/internal/databases"
)

type authHandler func(http.ResponseWriter, *http.Request, databases.User)

func (config *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid API Key")
			return
		} 

		user, err := config.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid API Key")
			return
		}

		handler(w, r, user)
 
	}
	
}