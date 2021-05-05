package auth

import (
	"context"
	"encoding/json"
	"fmt"
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

// updateSession will update the session containing the refresh token and then update its expiry
// date (if it's still available). If it's not available anymore, then it will return an error.
func updateSession(ctx context.Context, refreshToken string) (*db.SessionModel, error) {
	session, err := db.GetSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("the session does not exist")
	}

	if time.Now().After(session.ExpiryDate) {
		return nil, fmt.Errorf("the session has expired")
	}

	newExpiry := db.GetNewExpiryTime(session.Remember)
	session.ExpiryDate = newExpiry

	err = db.UpdateSessionByRefreshToken(ctx, refreshToken, session)
	if err != nil {
		return nil, fmt.Errorf("there was an error when updating the refresh token")
	}

	return session, nil
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

	res := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	json.NewEncoder(w).Encode(res)
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
		api.DefaultError(w, r, http.StatusBadRequest, api.DefaultErrorMessage)
		return
	}

	err = model.validateRegisterRequest(w, r, ctx)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, api.DefaultErrorMessage)
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

	res := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	json.NewEncoder(w).Encode(res)
}

// HandleRefresh will take a refresh token, check it against the db and then provide a new access
// token if it's acceptable.
func HandleRefresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	request, err := decodeRefreshRequest(w, r)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, api.DefaultErrorMessage)
		return
	}

	session, err := updateSession(ctx, request.RefreshToken)
	if err != nil {
		api.DefaultError(w, r, http.StatusBadRequest, api.DefaultErrorMessage)
		return
	}

	// generate jwt
	jwtPayload := newJwtPayload(session.UserID)
	accessToken, err := jwtPayload.convertToJwt()
	if err != nil {
		api.DefaultError(w, r, http.StatusInternalServerError, api.DefaultErrorMessage)
		return
	}

	res := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": session.RefreshToken,
	}

	json.NewEncoder(w).Encode(res)
}
