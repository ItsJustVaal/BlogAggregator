package main

import (
	"net/http"
)

func handleGetReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, basicResponse{
		Status: "ok",
	})
}

func handleGetErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
