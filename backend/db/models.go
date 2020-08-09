package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RaceSnippet is a model for a mongodb racesnippet
type RaceSnippet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Snippet string             `bson:"snippet" json:"snippet"`
}

// NewRaceSnippet returns a new RaceSnippet given a snippet string
func NewRaceSnippet(snippet string) *RaceSnippet {
	res := &RaceSnippet{
		Snippet: snippet,
	}

	return res
}

// User is a model for the a mongodb user
type User struct {
	ID           interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Email        string      `bson:"email" json:"email`
	Username     string      `bson:"username" json:"username"`
	Nickname     string      `bson:"nickname" json:"nickname"`
	Password     string      `bson:"password" json:"password"`
	Wpm          int         `bson:"wpm,minsize" json:"wpm"`
	RegisterDate time.Time   `bson:"register_date" json:"register_date"`
}

// NewUser returns a new User given all the details except for the ID
func NewUser(email string, username string, nickname string, password string, wpm int,
	registerDate time.Time) *User {

	res := &User{
		Email:        email,
		Username:     username,
		Nickname:     nickname,
		Password:     password,
		Wpm:          wpm,
		RegisterDate: registerDate,
	}
	return res
}
