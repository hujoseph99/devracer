package db

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNumDocuments(t *testing.T, collection *mongo.Collection) int64 {
	num, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		t.Fatal("Could not get the number of documents within collection")
	}
	return num
}

func TestAddAndDelete(t *testing.T) {
	testUser := NewUser("foo", "foo", "foo@foo.com", "", "", "", time.Now())

	collection := db.Database(DatabaseTypers).Collection(CollectionsUser)

	startingNum := getNumDocuments(t, collection)
	err := AddUser(context.Background(), testUser)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	id := testUser.ID.Hex()

	err = DeleteUserByID(context.Background(), id, RegularID)
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestAddAndDeleteOAuthUser(t *testing.T) {
	oauthID := "googleID"
	testUser := NewUser("foo", "foo", "foo@foo.com", oauthID, "", "", time.Now())

	collection := db.Database(DatabaseTypers).Collection(CollectionsUser)

	startingNum := getNumDocuments(t, collection)
	err := AddUser(context.Background(), testUser)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = DeleteUserByID(context.Background(), oauthID, GoogleID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestGetUserByID(t *testing.T) {
	testUser := NewUser("foo", "foo", "foo@foo.com", "", "", "", time.Now())

	err := AddUser(context.Background(), testUser)
	if err != nil {
		t.Fatal("Document was not added")
	}

	id := testUser.ID.Hex()

	foundUser, err := GetUserByID(context.Background(), id, RegularID)

	// checking username and password is good enough for me
	if err != nil || foundUser.ID.Hex() != id || foundUser.Username != testUser.Username ||
		foundUser.Password != testUser.Password {

		t.Fatal("Document was not found correctly")
	}

	err = DeleteUserByID(context.Background(), id, RegularID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestFindUserByUsername(t *testing.T) {
	testUser := NewUser("username", "foo", "foo@foo.com", "", "", "", time.Now())

	err := AddUser(context.Background(), testUser)
	if err != nil {
		t.Fatal("Document was not added")
	}

	id := testUser.ID.Hex()
	foundUser, err := GetUserByUsername(context.Background(), testUser.Username)

	// checking username and password is good enough for me
	if err != nil || foundUser.ID.Hex() != id || foundUser.Username != testUser.Username ||
		foundUser.Password != testUser.Password {

		t.Fatal("Document was not found correctly")
	}

	err = DeleteUserByID(context.Background(), id, RegularID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}

func TestFindUserByEmail(t *testing.T) {
	testUser := NewUser("username", "foo", "email@email.com", "", "", "", time.Now())

	err := AddUser(context.Background(), testUser)
	if err != nil {
		t.Fatal("Document was not added")
	}

	id := testUser.ID.Hex()
	foundUser, err := GetUserByEmail(context.Background(), testUser.Email)

	// checking username and password is good enough for me
	if err != nil || foundUser.ID.Hex() != id || foundUser.Username != testUser.Username ||
		foundUser.Password != testUser.Password || foundUser.Email != testUser.Email {

		t.Fatal("Document was not found correctly")
	}

	err = DeleteUserByID(context.Background(), id, RegularID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}
