package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProfileAddAndDelete(t *testing.T) {
	randID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the ID")
	}

	testProfile := NewProfile(randID, 100, 100, 100, 100, 100, 100)

	collection := client.client.Database(DatabaseTypers).Collection(CollectionProfiles)

	startingNum := getNumDocuments(t, collection)
	err = client.AddProfile(context.Background(), testProfile)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = client.DeleteProfileByID(context.Background(), testProfile.ID)
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}
