package config

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/hritik-hk/rss-aggregator/internal/auth"
	"github.com/hritik-hk/rss-aggregator/internal/database"
	"github.com/hritik-hk/rss-aggregator/utils"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DB *database.Queries
}

func NewDbConfig() (*dbConfig, error) {

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB connection url not found in env")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database: ", err)
	}

	dbQueries := database.New(connection)

	return &dbConfig{
		DB: dbQueries,
	}, nil

}

type authHandler func(http.ResponseWriter, *http.Request, database.User)

// define middlewareAuth method
func (dbConfig *dbConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIkey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := dbConfig.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			utils.RespondWithError(w, 404, "user not found")
			return
		}

		handler(w, r, user)
	}
}
