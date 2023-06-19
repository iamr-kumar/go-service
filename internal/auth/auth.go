package auth

import (
	"errors"
	"net/http"
	"strings"
)

/* extracts an API Key from header of a request
Example:
Authorization: Bearer <api_key> */
func GetApiKey(header http.Header) (string, error) {
	value := header.Get("Authorization")
	if len(value) < 8 {
		return "", errors.New("Invalid Authorization header")
	}

	vals := strings.Split(value, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid Authorization header")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("Invalid Authorization header")
	}

	return vals[1], nil
	
}