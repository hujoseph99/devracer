package typingMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToDB(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	// set client options -- get URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connet to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB")
}
