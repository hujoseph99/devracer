package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddAndDeleteSession(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	collection := db.Database(DatabaseTypers).Collection(CollectionsSessions)

	refreshToken := "1234123412341234"
	session := NewSession(refreshToken, userID, time.Now())

	startingNum := getNumDocuments(t, collection)
	err = AddSession(context.Background(), session)
	changedNum := getNumDocuments(t, collection)
	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = DeleteSessionByRefreshToken(context.Background(), refreshToken)
	changedNum = getNumDocuments(t, collection)
	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestGetSessionByRefreshToken(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	refreshToken := "1234123412341234"
	session := NewSession(refreshToken, userID, time.Now())
	err = AddSession(context.Background(), session)
	if err != nil {
		t.Fatal("Document was not added")
	}

	foundSession, err := GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not find the document")
	}

	if foundSession.RefreshToken != session.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		foundSession.ExpiryDate != session.ExpiryDate {
		t.Fatal("Did not find the same document")
	}

	err = DeleteSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestDeleteExistingSessionWhenAdding(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	firstRefreshToken := "1234123412341234"
	firstSession := NewSession(firstRefreshToken, userID, time.Now())

	err = AddSession(context.Background(), firstSession)
	if err != nil {
		t.Fatal("First session was not added")
	}

	secondRefreshToken := "1234123412341235"
	secondSession := NewSession(secondRefreshToken, userID, time.Now()) // same userid

	collection := db.Database(DatabaseTypers).Collection(CollectionsSessions)
	startingNum := getNumDocuments(t, collection)
	err = AddSession(context.Background(), secondSession)
	if err != nil {
		t.Fatal("Second session was not added")
	}
	changedNum := getNumDocuments(t, collection)
	if err != nil || changedNum != startingNum {
		t.Fatal("A session was not deleted")
	}

	foundSession, err := GetSessionByRefreshToken(context.Background(), secondRefreshToken)
	if err != nil {
		t.Fatal("Could not find the second session in the database")
	}
	if foundSession.RefreshToken != secondSession.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		foundSession.ExpiryDate != secondSession.ExpiryDate {
		t.Fatal("Second sesion is not the one in the database")
	}

	err = DeleteSessionByRefreshToken(context.Background(), secondRefreshToken)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestUpdateSesionByRefreshToken(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	refreshToken := "1234123412341234"
	session := NewSession(refreshToken, userID, time.Now())

	err = AddSession(context.Background(), session)
	if err != nil {
		t.Fatal("First session was not added")
	}

	newSession := NewSession(refreshToken, userID, time.Now().Add(time.Minute*15))
	err = UpdateSessionByRefreshToken(context.Background(), refreshToken, newSession)
	if err != nil {
		t.Fatal("Could not update the document")
	}

	foundSession, err := GetSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Could not find the updated document")
	}
	fmt.Println(session)
	fmt.Println(foundSession)
	if foundSession.RefreshToken != session.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		!foundSession.ExpiryDate.After(session.ExpiryDate.Add(time.Minute*14)) {
		t.Fatal("Session was not updated")
	}

	err = DeleteSessionByRefreshToken(context.Background(), refreshToken)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}
