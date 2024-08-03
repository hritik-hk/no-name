package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hritik-hk/rss-aggregator/internal/database"
	"github.com/hritik-hk/rss-aggregator/types"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response:%v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with code: %v, err: %v", code, msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func DatabaseUserToUser(dbUser database.User) types.User {
	return types.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIkey:    dbUser.ApiKey,
	}
}

func DatabaseFeedToFeed(dbFeed database.Feed) types.Feed {
	return types.Feed{
		ID:        dbFeed.ID,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
	}
}

func DatabaseFeedsToFeeds(feeds []database.Feed) []types.Feed {
	result := make([]types.Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = DatabaseFeedToFeed(feed)
	}
	return result
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) types.FeedFollow {
	return types.FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []types.FeedFollow {
	result := make([]types.FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}
