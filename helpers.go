package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func getAPIKey(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("No API key included")
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 {
		return "", errors.New("malformed authorization header")
	}
	return splitAuth[1], nil
}
