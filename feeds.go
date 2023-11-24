package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itsjustvaal/blogaggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	checker := jsonDecode{}
	err := decoder.Decode(&checker)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      checker.Name,
		Url:       checker.URL,
	})

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	finalFeed, FinalFFollow := databaseFeedToFeedResp(feed, feedFollow)
	respondWithJSON(w, http.StatusOK, CreateFeedResponse{
		Feed:       finalFeed,
		FeedFollow: FinalFFollow,
	})
}

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	var i []Feed
	for _, y := range feeds {
		i = append(i, Feed{
			ID:        y.ID,
			CreatedAt: y.CreatedAt,
			UpdatedAt: y.UpdatedAt,
			UserID:    y.UserID,
			Name:      y.Name,
			Url:       y.Url,
		})
	}
	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	checker := jsonDecode{}
	err := decoder.Decode(&checker)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    checker.Feed,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create feed follow")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFFollowToFFollow(feedFollow))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	checker := jsonDecode{}
	err := decoder.Decode(&checker)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	feedFollow, err := cfg.DB.GetFeedFollowByFeedID(r.Context(), checker.Feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Feed follow not found")
		return
	}

	if feedFollow.UserID != user.ID {
		respondWithError(w, http.StatusUnauthorized, "Cannot delete other users follows")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), checker.Feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, "Deleted Feed Follow")
}

func (cfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	var i []Feed
	for _, y := range feeds {
		if y.UserID == user.ID {
			i = append(i, Feed{
				ID:        y.ID,
				CreatedAt: y.CreatedAt,
				UpdatedAt: y.UpdatedAt,
				UserID:    y.UserID,
				Name:      y.Name,
				Url:       y.Url,
			})
		}
	}
	respondWithJSON(w, http.StatusOK, i)
}
