package auth

import (
	"context"
	"testing"
	"time"

	"github.com/hujoseph99/typing/backend/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateSession(t *testing.T) {
	userid, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not generate id")
	}

	refreshToken, err := createSession(context.Background(), userid, true)
	if err != nil {
		t.Fatal("Could not create session")
	}

	foundSession, err := db.GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not find the session")
	}

	if !foundSession.ExpiryDate.After(time.Now().Add(time.Hour*24*6)) ||
		foundSession.UserID.Hex() != userid.Hex() {
		t.Fatal("Did not find the correct session")
	}

	err = db.DeleteSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not delete the session")
	}
}

func TestCreateSessionNoRememberMe(t *testing.T) {
	userid, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not generate id")
	}

	refreshToken, err := createSession(context.Background(), userid, false)
	if err != nil {
		t.Fatal("Could not create session")
	}

	foundSession, err := db.GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not find the session")
	}

	if (!foundSession.ExpiryDate.After(time.Now().Add(time.Hour*23)) &&
		!foundSession.ExpiryDate.Before(time.Now().Add(time.Hour*25))) ||
		foundSession.UserID.Hex() != userid.Hex() {
		t.Fatal("Did not find the correct session")
	}

	err = db.DeleteSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not delete the session")
	}
}

func TestUpdateSession(t *testing.T) {
	userid, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not generate id")
	}

	refreshToken, err := createSession(context.Background(), userid, false)
	if err != nil {
		t.Fatal("Could not create session")
	}

	sessionInitial, err := db.GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not fetch the session")
	}

	newSession, err := updateSession(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("There was an error updating the session")
	}

	sessionFinal, err := db.GetSessionByRefreshToken(context.Background(), newSession.RefreshToken)
	if err != nil {
		t.Fatal("Could not fetch the updated session")
	}

	if sessionInitial.UserID.Hex() != sessionFinal.UserID.Hex() ||
		!sessionFinal.ExpiryDate.After(sessionInitial.ExpiryDate) ||
		(!sessionFinal.ExpiryDate.After(time.Now().Add(time.Hour*23)) &&
			!sessionFinal.ExpiryDate.Before(time.Now().Add(time.Hour*25))) {

		t.Fatal("Session was not updated correctly")
	}

	err = db.DeleteSessionByRefreshToken(context.Background(), sessionFinal.RefreshToken)
	if err != nil {
		t.Fatal("Could not delete the session")
	}
}

func TestUpdateSessionWithRememberMe(t *testing.T) {
	userid, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not generate id")
	}

	refreshToken, err := createSession(context.Background(), userid, true)
	if err != nil {
		t.Fatal("Could not create session")
	}

	sessionInitial, err := db.GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not fetch the session")
	}

	newSession, err := updateSession(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("There was an error updating the session")
	}

	sessionFinal, err := db.GetSessionByRefreshToken(context.Background(), newSession.RefreshToken)
	if err != nil {
		t.Fatal("Could not fetch the updated session")
	}

	if sessionInitial.UserID.Hex() != sessionFinal.UserID.Hex() ||
		!sessionFinal.ExpiryDate.After(sessionInitial.ExpiryDate) ||
		(!sessionFinal.ExpiryDate.After(time.Now().Add(time.Hour*24*6)) &&
			!sessionFinal.ExpiryDate.Before(time.Now().Add(time.Hour*24*8))) {

		t.Fatal("Session was not updated correctly")
	}

	err = db.DeleteSessionByRefreshToken(context.Background(), sessionFinal.RefreshToken)
	if err != nil {
		t.Fatal("Could not delete the session")
	}
}
