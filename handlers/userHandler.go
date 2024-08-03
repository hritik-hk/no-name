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

type UserHandler struct {
	DB *database.Queries
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	user, err := h.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaseUserToUser(user))

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))

}
