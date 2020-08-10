package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// GetPreferences returns the preferences for a user.
func (c *Client) GetPreferences(ctx context.Context, id string) (*PreferencesModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsPreferences)

	pref := PreferencesModel{}
	err := c.getDocumentFromCollectionByID(ctx, collection, id, &pref)

	if err != nil {
		return nil, err
	}
	return &pref, nil
}

// UpdatePreferences updates the preferences for a user.
func (c *Client) UpdatePreferences(ctx context.Context, id string, pref *PreferencesModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsPreferences)

	result := collection.FindOneAndUpdate(ctx, bson.M{"id": pref.ID}, bson.M{"$set": pref})

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

// AddPreferences adds a given pref to a mongo client.  Will return the error if there is an error.
func (c *Client) AddPreferences(ctx context.Context, pref *PreferencesModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	_, err := c.addDocumentToCollection(ctx, collection, pref)

	if err != nil {
		return err
	}

	return nil
}
