package db

import (
	"context"
)

// AddUser adds a given user to a mongo client.  If it is successful, then it
// will return the id in the form of a string and add it to the user model.  We
// are assuming that the password is already hashed and salted (if it is meant
// to be).
func (c *Client) AddUser(ctx context.Context, user *UserModel) (string, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	id, err := c.addDocumentToCollection(ctx, collection, user)

	if err != nil {
		return "", err
	}

	user.ID = id
	return id, nil
}

// DeleteUserByID will delete a user by the given id.  If it is successful, then
// it will return a nil error, otherwise it will return an error.
func (c *Client) DeleteUserByID(ctx context.Context, id string, idType int) error {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	err := c.deleteFromCollectionByID(ctx, collection, id, idType)
	return err
}

// FindUserByID finds a user given the id and then returns the user if it is
// successful.
func (c *Client) FindUserByID(ctx context.Context, id string) (*UserModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	var user UserModel
	err := c.getDocumentFromCollectionByID(ctx, collection, id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByUsername will find a user by their username in the db
func (c *Client) FindUserByUsername(ctx context.Context, username string) (*UserModel, error) {
	collection := c.client.Database(DatabaseTypers).Collection(CollectionsUser)

	params := make(map[string]string)
	params["username"] = username

	var user UserModel
	err := c.getDocumentFromCollection(ctx, collection, params, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

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
