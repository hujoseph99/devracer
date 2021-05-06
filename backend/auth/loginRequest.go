package auth

import (
	"encoding/json"
	"net/http"
)

// loginRequest is a model for user data that will be used for authentication
type loginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

func decodeLoginRequest(w http.ResponseWriter, r *http.Request) (*loginRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var model loginRequest
	err := decoder.Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
