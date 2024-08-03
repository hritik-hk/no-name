package handlers

import (
	"net/http"

	"github.com/hritik-hk/rss-aggregator/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithError(w, 400, "something went wrong!")
}
