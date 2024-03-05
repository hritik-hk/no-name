package main

import "net/http"

func handlerJSON(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
