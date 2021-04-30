package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddProfile adds a given profile to a mongo client.  If it is successful, then it
// will add the id to the given profile object and return a nil error.
func AddProfile(ctx context.Context, profile *ProfileModel) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionProfiles)

	id, err := addDocumentToCollection(ctx, collection, profile)
	if err != nil {
		return err
	}

	profile.ID = id
	return nil
}

// DeleteProfileByID will delete a profile by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func DeleteProfileByID(ctx context.Context, id primitive.ObjectID) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionProfiles)

	err := deleteFromCollectionByID(ctx, collection, id)
	if err != nil {
		return err
	}
	return nil
}

// GetProfileByID will get a profile from the database by the given id
func GetProfileByID(ctx context.Context, id primitive.ObjectID) (*ProfileModel, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionProfiles)

	var profileModel ProfileModel
	err := getDocumentFromCollectionByID(ctx, collection, id, &profileModel)
	if err != nil {
		return nil, err
	}

	return &profileModel, nil
}
