package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hujoseph99/typing/backend/common/api"
	"github.com/hujoseph99/typing/backend/common/utils"
	"github.com/hujoseph99/typing/backend/db"
)

type registerRequest struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func decodeRegisterRequest(w http.ResponseWriter, r *http.Request) (*registerRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var model registerRequest
	err := decoder.Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// validateUsername will validate the username. It will automatically add the response to the ResponseWriter
//	as well as return an error if there is a problem. Otherwise, the error will be nil.
func validateUsername(w http.ResponseWriter, r *http.Request, username string) error {
	if len(username) < 3 {
		api.DefaultError(w, r, http.StatusBadRequest, "The username must be at least 3 characters long")
		return fmt.Errorf("username was too short")
	}

	if utils.CheckValidUsernameCharacters(username) {
		api.DefaultError(w, r, http.StatusBadRequest, "The username cannot contain any special characters")
		return fmt.Errorf("username contained special characters")
	}

	return nil
}

// validatePassword will validate the password. It will automatically add the response to the ResponseWriter
// as well as return an error if there is a problem. Otherwise, the error will be nil.
func validatePassword(w http.ResponseWriter, r *http.Request, password string) error {
	if len(password) < 8 || len(password) > 64 {
		api.DefaultError(w, r, http.StatusBadRequest, "The password must be between 8 and 64 characters long.")
		return fmt.Errorf("username was too short")
	}

	// TODO: Maybe add this back in? Currently want to allow any characters
	// if utils.HasSpecialCharacters(username) {
	// 	api.DefaultError(w, r, http.StatusBadRequest, "The username cannot contain any special characters")
	// 	return fmt.Errorf("username contained special characters")
	// }

	return nil
}

// validateUsername will validate the username. It will automatically add the response to the ResponseWriter
//	as well as return an error if there is a problem. Otherwise, the error will be nil.
func validateNickname(w http.ResponseWriter, r *http.Request, nickname string) error {
	if len(nickname) <= 0 {
		api.DefaultError(w, r, http.StatusBadRequest, "The nickname must be at least 1 characters long")
		return fmt.Errorf("nickname was too short")
	}

	// Currently want to allow any character in the nickname

	return nil
}

// validateRegisterRequest will validate the given body of a register request. If there is an error, it will return the
// error, otherwise will return a nil error
func (model *registerRequest) validateRegisterRequest(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	// TODO: Fix this so that it does a case insensitive search
	existingUser, err := db.FindUserByUsername(ctx, model.Username)
	if err == nil || existingUser != nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The username is already in use. Please try again.")
		return fmt.Errorf("username was already in use")
	}

	// TODO: Fix this so that it does a case insensitive search
	existingUser, err = db.FindUserByEmail(ctx, model.Email)
	if err == nil || existingUser != nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The email is already in use. Please try again.")
		return fmt.Errorf("email was already in use")
	}

	err = validateUsername(w, r, model.Username)
	if err != nil {
		return err
	}

	err = validatePassword(w, r, model.Password)
	if err != nil {
		return err
	}

	err = validateNickname(w, r, model.Nickname)
	if err != nil {
		return err
	}

	fmt.Println(existingUser)
	return nil
}
