package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RaceHistoryModel is a model for a mongodb racesnippet
type RaceHistoryModel struct {
	ID               primitive.ObjectID     `bson:"_id,omitempty" json:"_id,omitempty"`
	SnippetID        string                 `bson:"snippetId" json:"snippetId"`
	RaceParticipants []RaceParticipantModel `bson:"raceParticipants" json:"raceParticipants"`
	Date             time.Time              `bson:"date" json:"date"`
}

// RaceParticipantModel is a model for a mongodb racesnippet
type RaceParticipantModel struct {
	PlayerID            string `bson:"playerId" json:"playerId"`
	Wpm                 int    `bson:"wpm" json:"wpm"`
	Time                int    `bson:"time" json:"time"`
	CharactersCorrect   int    `bson:"charactersCorrect" json:"charactersCorrect"`
	CharactersIncorrect int    `bson:"charactersIncorrect" json:"charactersIncorrect"`
	WordsIncorrect      int    `bson:"wordsIncorrect" json:"wordsIncorrect"`
}

// NewRaceHistoryModel returns a new NewRaceSnippet.
func NewRaceHistoryModel(snippetID string, raceParticipants []RaceParticipantModel, date *time.Time) *RaceHistoryModel {
	res := &RaceHistoryModel{
		SnippetID:        snippetID,
		RaceParticipants: raceParticipants,
		Date:             *date,
	}

	return res
}
