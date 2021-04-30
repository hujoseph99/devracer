package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetRaceHistoryByID returns a race history.
func GetRaceHistoryByID(ctx context.Context, id primitive.ObjectID) (*RaceHistoryModel, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	raceHistory := RaceHistoryModel{}
	err := getDocumentFromCollectionByID(ctx, collection, id, &raceHistory)

	if err != nil {
		return nil, err
	}
	return &raceHistory, nil
}

// UpdateRaceHistory updates a race history.
func UpdateRaceHistory(ctx context.Context, id primitive.ObjectID, raceHistory *RaceHistoryModel) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": raceHistory})

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

// AddRaceHistory adds a race history to a mongo client.  Will return the error if there is an error.
func AddRaceHistory(ctx context.Context, raceHistory *RaceHistoryModel) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	id, err := addDocumentToCollection(ctx, collection, raceHistory)

	if err != nil {
		return err
	}

	raceHistory.ID = id

	return nil
}

// DeleteRaceHistoryByID will delete a race history by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func DeleteRaceHistoryByID(ctx context.Context, id primitive.ObjectID) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionsRaceHistory)

	err := deleteFromCollectionByID(ctx, collection, id)
	return err
}
