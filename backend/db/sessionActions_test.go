package db

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddAndGetSession(t *testing.T) {
	userID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the userid")
	}

	sessionString := "1234123412341234"
	session := NewSession(sessionString, userID, time.Now())

}
