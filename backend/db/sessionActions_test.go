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

	sessionString := "1234123412341234"
	session := NewSession(sessionString, userID, time.Now())

	startingNum := getNumDocuments(t, collection)
	err = AddSession(context.Background(), session)
	changedNum := getNumDocuments(t, collection)
	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = DeleteSessionBySessionString(context.Background(), sessionString)
	changedNum = getNumDocuments(t, collection)
	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestGetSessionBySessionString(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	sessionString := "1234123412341234"
	session := NewSession(sessionString, userID, time.Now())
	err = AddSession(context.Background(), session)
	if err != nil {
		t.Fatal("Document was not added")
	}

	foundSession, err := GetSessionBySessionString(context.Background(), sessionString)
	if err != nil {
		t.Fatal("Could not find the document")
	}

	if foundSession.RefreshToken != session.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		foundSession.ExpiryDate != session.ExpiryDate {
		t.Fatal("Did not find the same document")
	}

	err = DeleteSessionBySessionString(context.Background(), sessionString)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestDeleteExistingSessionWhenAdding(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	firstSessionString := "1234123412341234"
	firstSession := NewSession(firstSessionString, userID, time.Now())

	err = AddSession(context.Background(), firstSession)
	if err != nil {
		t.Fatal("First session was not added")
	}

	secondSessionString := "1234123412341235"
	secondSession := NewSession(secondSessionString, userID, time.Now()) // same userid

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

	foundSession, err := GetSessionBySessionString(context.Background(), secondSessionString)
	if err != nil {
		t.Fatal("Could not find the second session in the database")
	}
	if foundSession.RefreshToken != secondSession.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		foundSession.ExpiryDate != secondSession.ExpiryDate {
		t.Fatal("Second sesion is not the one in the database")
	}

	err = DeleteSessionBySessionString(context.Background(), secondSessionString)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestUpdateSesionBySessionString(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	sessionString := "1234123412341234"
	session := NewSession(sessionString, userID, time.Now())

	err = AddSession(context.Background(), session)
	if err != nil {
		t.Fatal("First session was not added")
	}

	newSession := NewSession(sessionString, userID, time.Now().Add(time.Minute*15))
	err = UpdateSessionBySessionString(context.Background(), sessionString, newSession)
	if err != nil {
		t.Fatal("Could not update the document")
	}

	foundSession, err := GetSessionBySessionString(context.Background(), sessionString)
	if err != nil {
		t.Fatal("Could not find the updated document")
	}
	fmt.Println(session)
	fmt.Println(foundSession)
	if foundSession.RefreshToken != session.RefreshToken || foundSession.UserID.Hex() != userID.Hex() ||
		!foundSession.ExpiryDate.After(session.ExpiryDate.Add(time.Minute*14)) {
		t.Fatal("Session was not updated")
	}

	err = DeleteSessionBySessionString(context.Background(), sessionString)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}
