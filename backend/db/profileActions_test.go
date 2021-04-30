package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getTestProfile(t *testing.T) *ProfileModel {
	randID, err := primitive.ObjectIDFromHex("123412341234123412341234")
	if err != nil {
		t.Fatal("Could not create the ID")
	}

	testProfile := NewProfile(randID, 100, 100, 100, 100, 100, 100)

	return testProfile
}

func TestProfileAddAndDelete(t *testing.T) {
	testProfile := getTestProfile(t)

	collection := db.Database(DatabaseTypers).Collection(CollectionProfiles)

	startingNum := getNumDocuments(t, collection)
	err := AddProfile(context.Background(), testProfile)
	changedNum := getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum+1 {
		t.Fatal("Document was not added")
	}

	err = DeleteProfileByID(context.Background(), testProfile.ID)
	changedNum = getNumDocuments(t, collection)

	if err != nil || changedNum != startingNum {
		t.Fatal("Document was not deleted")
	}
}

func TestGetProfileByID(t *testing.T) {
	testProfile := getTestProfile(t)

	err := AddProfile(context.Background(), testProfile)
	if err != nil {
		t.Fatal("Document was not added")
	}

	foundProfile, err := GetProfileByID(context.Background(), testProfile.ID)

	// checking username and password is good enough for me
	if err != nil || foundProfile.ID.Hex() != testProfile.ID.Hex() ||
		foundProfile.RacesWon != testProfile.RacesWon ||
		foundProfile.RacesCompleted != testProfile.RacesCompleted {

		t.Fatal("Document was not found correctly")
	}

	err = DeleteProfileByID(context.Background(), testProfile.ID)
	if err != nil {
		t.Fatal("Document was not deleted")
	}
}
