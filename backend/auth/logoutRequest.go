package auth

import (
	"encoding/json"
	"net/http"
)

type logoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func decodeLogoutRequest(w http.ResponseWriter, r *http.Request) (*logoutRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var model logoutRequest
	err := decoder.Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
