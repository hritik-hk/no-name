package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hritik-hk/rss-aggregator/internal/database"
	"github.com/hritik-hk/rss-aggregator/utils"
)

type FeedFollowsHandler struct {
	DB *database.Queries
}

func (h *FeedFollowsHandler) CreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		FeedID uuid.UUID `json:"feedId"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't decode parameters: %v", err))
		return
	}

	feedFollow, err := h.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (h *FeedFollowsHandler) GetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := h.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't fetch feed follow")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (h *FeedFollowsHandler) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr) //converting from string to uuid type
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = h.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	deletedMsg := struct {
		Msg string `json:"msg"`
	}{
		Msg: "unfollowed successfully",
	}

	utils.RespondWithJSON(w, http.StatusOK, deletedMsg)
}
