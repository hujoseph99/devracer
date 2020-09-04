package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is a custom wrapper for the default mongo client
type Client struct {
	client *mongo.Client
}

// ConnectToDB connects to the database - this function also serves as the
// constructor for the Client
func ConnectToDB(ctx context.Context) (*Client, error) {
	// just in case
	if ctx == nil {
		ctx = context.Background()
	}

	// set client options -- get URI
	clientOptions := options.Client().ApplyURI(getMongoURI())

	// Connet to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB")
	newClient := &Client{
		client: client,
	}
	return newClient, nil
}
