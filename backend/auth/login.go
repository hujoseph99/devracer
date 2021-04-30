package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hujoseph99/typing/backend/common/api"
	"github.com/hujoseph99/typing/backend/db"
	"golang.org/x/crypto/bcrypt"
)

func decodeUserBody(w http.ResponseWriter, r *http.Request) (*db.UserModel, error) {
	decoder := json.NewDecoder(r.Body)
	var user db.UserModel
	err := decoder.Decode(&user)
	if err != nil {
		return nil, err
	} else if user.Username == "" || user.Password == "" {
		return nil, fmt.Errorf("invalid username or password")
	}
	return &user, nil
}

/*
 * Expects an object like:
 * {
 *  	username,
 *		password
 * }
 */
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")
	ctx := context.TODO()

	user, err := decodeUserBody(w, r)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, "Invalid request.")
		return
	}

	existingUser, err := db.FindUserByUsername(ctx, user.Username)
	if err != nil || existingUser == nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The username and password was not found. Please try again.")
		return
	}

	byteHash := []byte(existingUser.Password)
	bytePassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The username and password is incorrect. Please try again.")
	}

	jwtPayload := newJwtPayload(existingUser.ID)
	token, err := jwtPayload.convertToJwt()
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, "An error has occurred. Please try again.")
		return
	}

	jsonToken := map[string]string{
		"token": token,
	}

	json.NewEncoder(w).Encode(jsonToken)
}
