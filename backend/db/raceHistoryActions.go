package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetRaceHistoryByID returns a race history.
func (c *Client) GetRaceHistoryByID(ctx context.Context, id primitive.ObjectID) (*RaceHistoryModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	raceHistory := RaceHistoryModel{}
	err := c.getDocumentFromCollectionByID(ctx, collection, id, &raceHistory)

	if err != nil {
		return nil, err
	}
	return &raceHistory, nil
}

// UpdateRaceHistory updates a race history.
func (c *Client) UpdateRaceHistory(ctx context.Context, id primitive.ObjectID, raceHistory *RaceHistoryModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": raceHistory})

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

// AddRaceHistory adds a race history to a mongo client.  Will return the error if there is an error.
func (c *Client) AddRaceHistory(ctx context.Context, raceHistory *RaceHistoryModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	id, err := c.addDocumentToCollection(ctx, collection, raceHistory)

	if err != nil {
		return err
	}

	raceHistory.ID = id

	return nil
}

// DeleteRaceHistoryByID will delete a race history by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func (c *Client) DeleteRaceHistoryByID(ctx context.Context, id primitive.ObjectID) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	err := c.deleteFromCollectionByID(ctx, collection, id)
	return err
}
