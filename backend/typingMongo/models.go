package typingMongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RaceSnippet struct {
	Snippet string `bson:"racesnippet", json:"racesnippet"`
}

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty", json:"id,omitempty"`
	Username     string             `bson:"username", json:"username"`
	Nickname     string             `bson:"nickname", json:"nickname"`
	Password     string             `bson:"password", json:"password"`
	Wpm          int                `bson:"wpm", json:"wpm"`
	RegisterDate time.Time          `bson:"register_date", json:"register_date"`
}
