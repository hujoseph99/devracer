package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// GetPreferences returns the preferences for a user.
func (c *Client) GetPreferences(ctx context.Context, id string) (*PreferencesModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsPreferences)

	var pref PreferencesModel
	err := c.getDocumentFromCollectionByID(ctx, collection, id, &pref)

	if err != nil {
		return nil, err
	}
	return &pref, nil
}

// UpdatePreferences updates the preferences for a user.
func (c *Client) UpdatePreferences(ctx context.Context, id string, pref *PreferencesModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsPreferences)

	result := collection.FindOneAndUpdate(ctx, bson.M{"id": pref.ID}, pref)

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
