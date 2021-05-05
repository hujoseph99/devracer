package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hujoseph99/typing/backend/common/api"
	"github.com/hujoseph99/typing/backend/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// createSession creates a session for the given userid and then return thes refreshToken string
// if successful. Otherwise, an error will be returned
func createSession(ctx context.Context, userid primitive.ObjectID, rememberMe bool) (string, error) {
	// create session
	refreshToken := db.GenerateRefreshToken()
	expiryTime := time.Now()
	if rememberMe {
		expiryTime = expiryTime.Add(db.RememberMeExpiryTime)
	} else {
		expiryTime = expiryTime.Add(db.DefaultExpiryTime)
	}
	session := db.NewSession(refreshToken, userid, expiryTime, rememberMe)
	err := db.AddSession(ctx, session)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

/*
 * Expects an object like:
 * {
 *  	username,
 *		password,
 *    rememberMe
 * }
 */
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")
	ctx := context.TODO()

	login, err := decodeLoginRequest(w, r)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, "Invalid request.")
		return
	}

	existingUser, err := db.FindUserByUsername(ctx, login.Username)
	if err != nil || existingUser == nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The username and password was not found. Please try again.")
		return
	}

	byteHash := []byte(existingUser.Password)
	bytePassword := []byte(login.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		api.DefaultError(w, r, http.StatusUnauthorized, "The username and password was not found. Please try again.")
	}

	// generate jwt
	jwtPayload := newJwtPayload(existingUser.ID)
	accessToken, err := jwtPayload.convertToJwt()
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	// create session
	refreshToken, err := createSession(ctx, existingUser.ID, login.RememberMe)
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	jsonToken := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	json.NewEncoder(w).Encode(jsonToken)
}

/*
 * Expects an object like:
 * {
 *  	username,
 *		nickname,
 *		password,
 *		email
 * }
 */
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	model, err := decodeRegisterRequest(w, r)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, "Invalid request")
	}

	err = model.validateRegisterRequest(w, r, ctx)
	if err != nil {
		return
	}

	// hash the password
	bytePassword := []byte(model.Password)
	hashed, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	newUser := db.NewUser(model.Username, string(hashed), model.Email, "", "", "", time.Now())
	newPreferences := db.NewPreferences(newUser.ID, model.Nickname)
	newProfile := db.NewProfile(newUser.ID, 0, 0, 0, 0, 0, 0)

	err = db.RegisterUser(ctx, newUser, newProfile, newPreferences)
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	jwtPayload := newJwtPayload(newUser.ID)
	accessToken, err := jwtPayload.convertToJwt()
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	// create session
	refreshToken, err := createSession(ctx, newUser.ID, false)
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	jsonToken := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	json.NewEncoder(w).Encode(jsonToken)
}

// HandleRefresh will take a refresh token, check it against the db and then provide a new access
// token if it's acceptable.
func HandleRefresh(w http.ResponseWriter, r *http.Request) {

}
