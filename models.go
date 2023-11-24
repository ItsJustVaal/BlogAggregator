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
	Name string    `json:"name"`
	URL  string    `json:"url"`
	Feed uuid.UUID `json:"feed_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apikey    string    `json:"apikey"`
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

type CreateFeedResponse struct {
	Feed       Feed       `json:"feed"`
	FeedFollow FeedFollow `json:"feed_follow"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Post struct {
	ID           uuid.UUID `json:"id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	Description  string    `json:"description"`
	Published_at time.Time `json:"published_at"`
	Feed_id      uuid.UUID `json:"feed_id"`
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

func databaseFeedToFeedResp(feed database.Feed, feedFollow database.FeedFollow) (Feed, FeedFollow) {
	return Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		}, FeedFollow{
			ID:        feedFollow.ID,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feedFollow.UpdatedAt,
			FeedID:    feedFollow.FeedID,
			UserID:    feedFollow.UserID,
		}
}

func databaseFFollowToFFollow(feed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		FeedID:    feed.FeedID,
		UserID:    feed.UserID,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	var finalPosts []Post
	for _, x := range posts {
		finalPosts = append(finalPosts, Post{
			ID:           x.ID,
			Created_at:   x.CreatedAt,
			Updated_at:   x.UpdatedAt,
			Title:        x.Title,
			Url:          x.Url,
			Description:  x.Description.String,
			Published_at: x.PublishedAt.Time,
			Feed_id:      x.FeedID,
		})
	}
	return finalPosts
}
