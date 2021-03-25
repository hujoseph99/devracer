package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

// AddUser adds a given user to a mongo client.  If it is successful, then it
// will add it to the given user object and return a nil error.  We are assuming
// that the password is already hashed and salted.
func (c *Client) AddUser(ctx context.Context, user *UserModel) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	id, err := c.addDocumentToCollection(ctx, collection, user)

	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// DeleteUserByID will delete a user by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func (c *Client) DeleteUserByID(ctx context.Context, id string, idType int) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	bsonID, err := getBsonID(id, idType)
	if err != nil {
		return err
	}

	del, err := collection.DeleteOne(ctx, bsonID)
	if err != nil {
		return err
	} else if del.DeletedCount == 0 {
		return fmt.Errorf("ERROR: The document was not found")
	}

	return nil
}

// GetUserByID finds a user given the id and then returns the user if it is
// successful.
func (c *Client) GetUserByID(ctx context.Context, id string, idType int) (*UserModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	bsonID, err := getBsonID(id, idType)
	if err != nil {
		return nil, err
	}

	var user UserModel
	_ = collection.FindOne(ctx, bsonID).Decode(&user)

	return &user, nil
}

// FindUserByUsername will find a user by their username in the db
// func (c *Client) FindUserByUsername(ctx context.Context, username string) (*UserModel, error) {
// 	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

// 	params := make(map[string]string)
// 	params["username"] = username

// 	var user UserModel
// 	err := c.getDocumentFromCollection(ctx, collection, params, &user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// // FindUserByEmail will find a user by their email in the db
// func (c *Client) FindUserByEmail(ctx context.Context, email string) (*UserModel, error) {
// 	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

// 	params := make(map[string]string)
// 	params["email"] = email

// 	var user UserModel
// 	err := c.getDocumentFromCollection(ctx, collection, params, &user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
