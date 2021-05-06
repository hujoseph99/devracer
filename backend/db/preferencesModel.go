package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PreferencesModel is a model for storing preferences
type PreferencesModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DisplayName string             `bson:"displayName" json:"displayName"`
}

// NewPreferences is the constructor for PreferencesModel - it creates a new PreferencesModel object
func NewPreferences(id primitive.ObjectID, displayName string) *PreferencesModel {
	res := &PreferencesModel{
		ID:          id,
		DisplayName: displayName,
	}
	return res
}
