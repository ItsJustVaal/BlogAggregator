package main

import "github.com/itsjustvaal/blogaggregator/internal/database"

type apiConfig struct {
	DB *database.Queries
}

type basicResponse struct {
	Status string `json:"status"`
}

type errorResp struct {
	Error string `json:"error"`
}

type jsonDecode struct {
	Body string `json:"body"`
}
