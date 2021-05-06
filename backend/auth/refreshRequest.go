package auth

import (
	"encoding/json"
	"net/http"
)

type refreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func decodeRefreshRequest(w http.ResponseWriter, r *http.Request) (*refreshRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var model refreshRequest
	err := decoder.Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
