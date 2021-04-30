package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// addDocumentToCollection will add a document to the given collection.  If it is
// then it will return the id in the form of a string.
func addDocumentToCollection(ctx context.Context,
	collection *mongo.Collection, doc interface{}) (primitive.ObjectID, error) {

	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// should not ever fail because gotten directly from insert command
	id, _ := insertResult.InsertedID.(primitive.ObjectID)

	log.Printf("Inserted: %v\n", id)
	return id, nil
}

// deleteFromCollectionByID will delete a document from the given collection by ID.
// If it is successful, then the error will be nil, otherwise it will be a valid
// error.
func deleteFromCollectionByID(ctx context.Context,
	collection *mongo.Collection, id primitive.ObjectID) error {

	del, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if del.DeletedCount == 0 {
		return fmt.Errorf("ERROR: Could not delete the document")
	}

	log.Printf("Deleted object with ID: %v\n", id)
	return nil
}

// getDocumentFromCollectionByID will get a document from a given collection
// matches the given id and will populate it into the given model.
func getDocumentFromCollectionByID(ctx context.Context,
	collection *mongo.Collection, id primitive.ObjectID, model interface{}) error {

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(model)
	if err != nil {
		return err
	}
	return nil
}

// getDocumentFromCollection will get a document from a given collection that
// matches any of the given params if first param indicates the field name
// and the second param indicates the value of the param that is being searched for.
// The document will be populated into the given model (if it did not error).
func getDocumentFromCollection(ctx context.Context,
	collection *mongo.Collection, params map[string]string,
	model interface{}) error {

	matches := make([]bson.M, 0, len(params))
	for key, val := range params {
		matches = append(matches, bson.M{key: bson.M{"$eq": val}})
	}
	query := bson.M{"$or": matches}

	err := collection.FindOne(ctx, query).Decode(model)
	return err
}
