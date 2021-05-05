package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/hujoseph99/typing/backend/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
