package models

import (
	"github.com/hujoseph99/typing/backend/db"
)

type NewPracticeRace struct {
	Snippet   *db.Snippet
	TimeLimit int
}

// GetTimeLimit
func GetTimeLimit(snippet *db.Snippet) int {
	return len(snippet.RaceContent)
}

// NewNewPracticeRace returns a NewPracticeRace with the given snippet and
//	timeLimit
func NewNewPracticeRace(snippet *db.Snippet, timeLimit int) *NewPracticeRace {
	return &NewPracticeRace{
		Snippet:   snippet,
		TimeLimit: timeLimit,
	}
}
