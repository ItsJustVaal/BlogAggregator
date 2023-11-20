package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/itsjustvaal/blogaggregator/internal/database"
)

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

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apikey    string    `json:"apikey"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Apikey:    user.Apikey,
	}
}
