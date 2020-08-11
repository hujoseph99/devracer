package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PreferencesModel is a model for storing preferences
type PreferencesModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	DisplayName string             `bson:"displayName" json:"displayName"`
	Email       string             `bson:"email" json:"email"`
	Theme       string             `bson:"theme" json:"theme"`
}

// NewPreferences is the constructor for PreferencesModel - it creates a new PreferencesModel object
func NewPreferences(id primitive.ObjectID, displayName string, email string, theme string) *PreferencesModel {

	res := &PreferencesModel{
		ID:          id,
		DisplayName: displayName,
		Email:       email,
		Theme:       theme,
	}
	return res
}

// Equals returns true if the values in the struct are equal.
func (p *PreferencesModel) Equals(com *PreferencesModel) bool {

	if p.ID.String() == com.ID.String() &&
		p.DisplayName == com.DisplayName &&
		p.Email == com.Email &&
		p.Theme == com.Theme {
		return true
	}
	return false

}
