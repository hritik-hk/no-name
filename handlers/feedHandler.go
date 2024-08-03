package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hritik-hk/rss-aggregator/internal/database"
	"github.com/hritik-hk/rss-aggregator/utils"
)

type FeedHandler struct {
	DB *database.Queries
}

func (h *FeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	feed, err := h.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaseFeedToFeed(feed))

}

func (h *FeedHandler) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't fetch feeds")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedsToFeeds(feeds))
}
