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
	collection *mongo.Collection, doc interface{}) (string, error) {

	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	// should not ever fail because gotten directly from insert command
	id, _ := insertResult.InsertedID.(primitive.ObjectID)
	res := id.Hex()

	log.Printf("Inserted: %v\n", id)
	return res, nil
}

// getBsonId gets a bson.M object for
func getBsonId(id string, idType int) (bson.M, error) {
	var idKey string
	var idValue interface{}

	if idType == RegularID {
		idKey = "_id"
	} else if idType == GoogleID {
		idKey = "googleID"
	} else if idType == GithubID {
		idKey = "githubID"
	} else {
		idKey = "facebookID"
	}

	if idType == RegularID {
		objID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return nil, err
		}

		idValue = objID
	} else {
		idValue = idKey
	}

	// TODO: have to write a test to see if this works -- might not work for regularId
	// cause it requires an primite.objectId
	return bson.M{idKey: idValue}, nil

}

// deleteFromCollectionByID will delete a document from the given collection by ID.
// If it is successful, then the error will be nil, otherwise it will be a valid
// error.
func (c *Client) deleteFromCollectionByID(ctx context.Context,
	collection *mongo.Collection, id string, idType int) error {

	// objID, err := primitive.ObjectIDFromHex(id)

	// if err != nil {
	// 	return err
	// }

	bsonID, err := getBsonId(id, idType)
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
	collection *mongo.Collection, id string, model interface{}) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(model)
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
