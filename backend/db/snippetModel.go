package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Emulated enum for the different languages that we support
const (
	LanguageEnglish    = iota // 0
	LanguageCPP        = iota // 1
	LanguageGo         = iota // 2
	LanguageJavascript = iota // 3
	LanguagePython     = iota // 4
)

// Snippet is a model for a mongodb
type Snippet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RaceContent string             `bson:"snippet"`
	TokenCount  int                `bson:"-"`
	Language    int                `bson:"language"`
	DateCreated time.Time          `bson:"dateCreated"`
	// UsageCount  int                `bson:"usageCount" json:"usageCount"`
}

// NewSnippet is the constructor for Snippet - it creates a snippet object
func NewSnippet(raceContent string, language int, dateCreated time.Time) *Snippet {
	res := &Snippet{
		RaceContent: raceContent,
		TokenCount:  len(raceContent),
		Language:    language,
		DateCreated: dateCreated,
	}

	return res
}
