package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddRaceSnippet adds a given RaceSnippet to a mongo client.  If it is successful
// then it will return the id in the form of a string
func (c *Client) AddRaceSnippet(ctx context.Context, snippet *RaceSnippet) (string, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceSnippets)

	id, err := c.addDocumentToCollection(ctx, collection, snippet)
	return id, err
}

// DeleteRaceSnippetByID will delete a race snippet by the id given.  If it is successful,
// then it will return a nil error, otherwise it will return an error.
func (c *Client) DeleteRaceSnippetByID(ctx context.Context, id primitive.ObjectID) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceSnippets)

	err := c.deleteFromCollectionByID(ctx, collection, id)

	return err
}

// GetRaceSnippetByID gets a race snippet by ID and then returns the RaceSnippet if it is
// successful.
func (c *Client) GetRaceSnippetByID(ctx context.Context, id primitive.ObjectID) (*RaceSnippet, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceSnippets)

	var raceSnippet RaceSnippet
	err := c.getDocumentFromCollectionByID(ctx, collection, id, &raceSnippet)
	if err != nil {
		return nil, err
	}
	return &raceSnippet, nil
}

// GetRandomRaceSnippet gets a random race snippet if one can be found.  Otherwise,
// it will return an error.
func (c *Client) GetRandomRaceSnippet(ctx context.Context) (*RaceSnippet, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsRaceSnippets)
	randSnippet := bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{randSnippet}, opts)
	if err != nil {
		return nil, err
	}

	var res []bson.M
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	} else if len(res) == 0 {
		return nil, fmt.Errorf("No races were found")
	}
	var snippet RaceSnippet
	bsonBytes, _ := bson.Marshal(res[0])
	bson.Unmarshal(bsonBytes, &snippet)
	return &snippet, nil
}
