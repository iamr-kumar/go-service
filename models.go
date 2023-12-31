package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/iamr-kumar/go-service/internal/databases"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Email string `json:"email"`
	ApiKey string `json:"api_key"`
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
}

func convertDbUserToUser(dbUser databases.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		Email: dbUser.Email,
		ApiKey: dbUser.ApiKey,
	}
}

func convertDbFeedToFeed(dbFeed databases.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func convertDbFeedsToFeed(dbFeeds []databases.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, convertDbFeedToFeed(dbFeed))
	}

	return feeds
}

func convertDbFeedFollowToFeedFollow(dbFeedFollow databases.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID: dbFeedFollow.FeedID,
		UserID: dbFeedFollow.UserID,
	}
}

func convertDbFeedFollowsToFeedFollows(dbFeedFollows []databases.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, convertDbFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedFollows
}