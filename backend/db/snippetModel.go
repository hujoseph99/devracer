package db

import "go.mongodb.org/mongo-driver/bson/primitive"

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
