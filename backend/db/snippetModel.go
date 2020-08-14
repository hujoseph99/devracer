package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Emulated enum for the different languages that we support
const (
	LanguageCPP        = iota
	LanguageJavascript = iota
	LanguagePython     = iota
	LanguageGo         = iota
)

// Snippet is a model for a mongodb
type Snippet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Snippet     string             `bson:"snippet" json:"snippet"`
	Language    int                `bson:"language" json:"snippetId"`
	UsageCount  int                `bson:"usageCount" json:"usageCount"`
	DateCreated time.Time          `bson:"dateCreated" json:"dateCreated"`
}

// NewSnippet is the constructor for Snippet - it creates a snippet object
func NewSnippet(snippet string, language int, usageCount int,
	dateCreated time.Time) *Snippet {

	res := &Snippet{
		Snippet:     snippet,
		Language:    language,
		UsageCount:  usageCount,
		DateCreated: dateCreated,
	}

	return res
}
