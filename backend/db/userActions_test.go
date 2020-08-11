package db

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *Client

func getNumDocuments(t *testing.T, collection *mongo.Collection) int64 {
	num, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		t.Fatal("Could not get the number of documents within collection")
	}
	return num
}

func TestAddAndDelete(t *testing.T) {
	testUser := NewUser("foo", "foo", "foo", "", "", "", time.Now())

	collection := client.client.Database(DatabaseTypers).Collection(CollectionsUser)

	startingNum := getNumDocuments(t, collection)
	id, err := client.AddUser(context.Background(), testUser)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = client.DeleteUserByID(context.Background(), id, RegularID)
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestAddAndDeleteOAuthUser(t *testing.T) {
	oauthID := "googleID"
	testUser := NewUser("foo", "foo", "foo", oauthID, "", "", time.Now())

	collection := client.client.Database(DatabaseTypers).Collection(CollectionsUser)

	startingNum := getNumDocuments(t, collection)
	_, err := client.AddUser(context.Background(), testUser)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = client.DeleteUserByID(context.Background(), oauthID, GoogleID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestFindUserByID(t *testing.T) {

}

func TestMain(m *testing.M) {
	var err error
	client, err = ConnectToDB(context.Background())

	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
