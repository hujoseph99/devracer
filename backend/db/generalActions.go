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
func (c *Client) addDocumentToCollection(ctx context.Context,
	collection *mongo.Collection, doc interface{}) (*primitive.ObjectID, error) {

	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	// should not ever fail because gotten directly from insert command
	id, _ := insertResult.InsertedID.(primitive.ObjectID)

	log.Printf("Inserted: %v\n", id)
	return &id, nil
}

// getBsonId gets a bson.M object for
func getBsonID(id string, idType int) (bson.M, error) {
	var idKey string

	if idType == RegularID {
		idKey = "_id"
	} else if idType == GoogleID {
		idKey = "googleID"
	} else if idType == GithubID {
		idKey = "githubID"
	} else if idType == FacebookID {
		idKey = "facebookID"
	}

	if idType == RegularID {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		return bson.M{idKey: objID}, nil
	}
	return bson.M{idKey: id}, nil
}

// deleteFromCollectionByID will delete a document from the given collection by ID.
// If it is successful, then the error will be nil, otherwise it will be a valid
// error.
func (c *Client) deleteFromCollectionByID(ctx context.Context,
	collection *mongo.Collection, id string, idType int) error {

	bsonID, err := getBsonID(id, idType)
	if err != nil {
		return err
	}

	del, err := collection.DeleteOne(ctx, bsonID)
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
func (c *Client) getDocumentFromCollectionByID(ctx context.Context,
	collection *mongo.Collection, id string, idType int, model interface{}) error {

	bsonID, err := getBsonID(id, idType)
	if err != nil {
		return err
	}

	err = collection.FindOne(ctx, bsonID).Decode(model)
	if err != nil {
		return err
	}
	return nil
}

// getDocumentFromCollection will get a document from a given collection that
// matches any of the given params if first param indicates the field name
// and the second param indicates the value of the param that is being searched for.
// The document will be populated into the given model (if it did not error).
func (c *Client) getDocumentFromCollection(ctx context.Context,
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
