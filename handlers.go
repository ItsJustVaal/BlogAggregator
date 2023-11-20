package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itsjustvaal/blogaggregator/internal/database"
)

func handleGetReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, basicResponse{
		Status: "ok",
	})
}

func handleGetErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	checker := jsonDecode{}
	err := decoder.Decode(&checker)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      checker.Body,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := getAPIKey(r.Header)
	log.Printf("APIKEY FOUND, %s\n", apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "No API Key Found")
		return
	}

	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
