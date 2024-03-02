package main

import "net/http"

func handleJSON(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
