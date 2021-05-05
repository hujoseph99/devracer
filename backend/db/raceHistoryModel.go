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

// NewRaceHistory returns a new RaceHistoryModel.
func NewRaceHistory(snippetID string, raceParticipants []RaceParticipantModel, date time.Time) *RaceHistoryModel {
	res := &RaceHistoryModel{
		SnippetID:        snippetID,
		RaceParticipants: raceParticipants,
		Date:             date.UTC().Round(time.Millisecond), // mongodb converts to UTC, make it consistent with our models
	}

	return res
}

// NewRaceParticipant returns a new NewRaceSnippet.
func NewRaceParticipant(playerID string, wpm int, time int, charactersCorrect int, charactersIncorrect int, wordsIncorrect int) *RaceParticipantModel {
	res := &RaceParticipantModel{
		PlayerID:            playerID,
		Wpm:                 wpm,
		Time:                time,
		CharactersCorrect:   charactersCorrect,
		CharactersIncorrect: charactersIncorrect,
		WordsIncorrect:      wordsIncorrect,
	}

	return res
}
