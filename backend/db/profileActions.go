package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddProfile adds a given profile to a mongo client.  If it is successful, then it
// will add the id to the given profile object and return a nil error.
func (c *Client) AddProfile(ctx context.Context, profile *ProfileModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionProfiles)

	id, err := c.addDocumentToCollection(ctx, collection, profile)
	if err != nil {
		return err
	}

	profileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	profile.ID = profileID
	return nil
}

// DeleteProfileByID will delete a profile by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func (c *Client) DeleteProfileByID(ctx context.Context, id primitive.ObjectID) error {

	collection := c.client.Database(DatabaseTypers).Collection(CollectionProfiles)

	err := c.deleteFromCollectionByID(ctx, collection, id)
	if err != nil {
		return err
	}
	return nil
}
