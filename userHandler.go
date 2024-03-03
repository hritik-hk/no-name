package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hritik-hk/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))

}
