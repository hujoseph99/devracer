package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hujoseph99/typing/backend/db"
)

// getRanomRaceSnippet will get a random race snippet from the db and return
// it to the client in JSON format
func getRandomSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	snippet, err := db.GetRandomSnippet(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(snippet); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
