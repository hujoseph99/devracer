package api

import (
	"context"
	"encoding/json"
	"net/http"
)

// getRanomRaceSnippet will get a random race snippet from the db and return
// it to the client in JSON format
func (api *API) getRandomSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	snippet, err := api.Database.GetRandomSnippet(context.TODO())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(snippet); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
