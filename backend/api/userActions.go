package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hujoseph99/typingBackend/db"
	"golang.org/x/crypto/bcrypt"
)

const defaultErrorMessage = "An error has occurred. Please try again."

// hashAndSalt encrypts a password for us
func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// defaultError is a default error function that will write an error message to
// the response
func defaultError(w http.ResponseWriter, r *http.Request,
	status int, message string) {

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newErrorMessage(message))
}

// decodoeUserBody will decode the body of the user request into a User object
// (if it can), otherwise it will return an error
func decodeUserBody(w http.ResponseWriter, r *http.Request) (*db.User, error) {
	decoder := json.NewDecoder(r.Body)
	var user db.User
	err := decoder.Decode(&user)
	if err != nil {
		return nil, err
	} else if user.Username == "" || user.Password == "" {
		return nil, fmt.Errorf("Invalid username or password")
	}
	return &user, nil
}

// returnUserToClient will return a user to the client
func returnUserToClient(w http.ResponseWriter, r *http.Request, user *db.User) {
	// user to return to client
	retUser, err := NewUserReturnToClient(user)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, defaultErrorMessage)
		return
	}

	// map user to claims
	token, err := retUser.convertToJwt()
	if err != nil {
		defaultError(w, r, http.StatusInternalServerError, defaultErrorMessage)
		return
	}

	// prepare jsonToken to be returned to user in JSON format
	jsonToken := map[string]string{
		"token": token,
	}

	json.NewEncoder(w).Encode(jsonToken)
}

// registerUser will register a user through an http request.
// We assume that we will get a request with a JSON body that will contain
// an email, username, nickname, and password
func (myAPI *API) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	user, err := decodeUserBody(w, r)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, "Invalid request.")
		return
	}

	// check if username already exists
	existingUser, err := myAPI.Database.FindUserByUsername(ctx, user.Username)
	if err == nil || existingUser != nil {
		defaultError(w, r, http.StatusUnauthorized, "The username is already in use. Please try again.")
		return
	}

	// Check if the email already exists
	existingUser, err = myAPI.Database.FindUserByEmail(ctx, user.Email)
	if err == nil || existingUser != nil {
		defaultError(w, r, http.StatusUnauthorized, "The email is already in use. Please try again.")
		return
	}

	// Fill in the rest of the fields with default data
	user.Wpm = 0
	user.RegisterDate = time.Now()

	// hash and salt the password
	bytePassword := []byte(user.Password)
	hashed, err := hashAndSalt(bytePassword)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, defaultErrorMessage)
		return
	}
	user.Password = hashed

	// add user to db
	err = myAPI.Database.AddUser(ctx, user)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, defaultErrorMessage)
		return
	}

	returnUserToClient(w, r, user)
}

// loginUser will log in a user
// We assume that we will get a request with a JSON body that will contain
// a username, and password.
func (myAPI *API) loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	user, err := decodeUserBody(w, r)
	if err != nil {
		defaultError(w, r, http.StatusBadRequest, "Invalid request.")
		return
	}

	// check if username already exists
	existingUser, err := myAPI.Database.FindUserByUsername(ctx, user.Username)
	if err != nil || existingUser == nil {
		defaultError(w, r, http.StatusUnauthorized, "The username was not found. Please try again.")
		return
	}

	byteHash := []byte(existingUser.Password)
	bytePassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		defaultError(w, r, http.StatusUnauthorized, "The password is incorrect. Please try again.")
		return
	}

	returnUserToClient(w, r, existingUser)
}
