package main

import "net/http"

func handleError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusNotFound, "Something went wrong")
}