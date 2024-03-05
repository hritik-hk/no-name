package auth

import (
	"errors"
	"net/http"
	"strings"
)

// APIkey {user_api_key}
func GetAPIkey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}

	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) != 2 || splitHeader[0] != "APIkey" {
		return "", errors.New("malformed authorization header")
	}

	return splitHeader[1], nil
}
