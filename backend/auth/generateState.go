package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

const stateLength = 16
const oauthCookieName = "OAuthState"

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// generateEcryptedOAuthStateString will geneate a state string
//	that can be passed along OAuth requests
func generateStateOAuthCookie(w http.ResponseWriter) (string, error) {
	var expiration = time.Now().Add(10 * time.Minute)

	stateBytes, err := generateRandomBytes(stateLength)
	if err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(stateBytes)
	cookie := http.Cookie{
		Name:     oauthCookieName,
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state, nil
}
