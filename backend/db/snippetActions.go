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

// AddSnippet adds a given snippet to a mongo client.  If it is successful
// then it will add the id to the given snippet.  Otherwise, it will return an
// error.
func AddSnippet(ctx context.Context, snippet *Snippet) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSnippets)

	id, err := addDocumentToCollection(ctx, collection, snippet)
	if err != nil {
		return err
	}

	snippet.ID = id
	return nil
}

// DeleteSnippetByID will delete a snippet by the id given.  If it is successful,
// then it will return a nil error, otherwise it will return an error.
func DeleteSnippetByID(ctx context.Context, id primitive.ObjectID) error {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSnippets)

	return deleteFromCollectionByID(ctx, collection, id)
}

// GetSnippetByID gets a race snippet by ID and then returns the RaceSnippet if it is
// successful.
func GetSnippetByID(ctx context.Context, id primitive.ObjectID) (*Snippet, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSnippets)

	var snippet Snippet
	err := getDocumentFromCollectionByID(ctx, collection, id, &snippet)
	if err != nil {
		return nil, err
	}
	return &snippet, nil
}

// GetRandomSnippet gets a random race snippet if one can be found.  Otherwise,
// it will return an error.
func GetRandomSnippet(ctx context.Context) (*Snippet, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSnippets)

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
		return nil, fmt.Errorf("no races were found")
	}
	var snippet Snippet
	bsonBytes, _ := bson.Marshal(res[0])
	bson.Unmarshal(bsonBytes, &snippet)

	snippet.TokenCount = len(snippet.RaceContent)
	return &snippet, nil
}
