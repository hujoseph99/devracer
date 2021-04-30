package db

import (
	"context"
	"log"

	"github.com/hujoseph99/typing/backend/secret"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Client
)

func InitDatabase() {
	clientOptions := options.Client().ApplyURI(secret.MongoURI)

	var err error
	db, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = db.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to DB")
}
