package models

import (
	"github.com/hujoseph99/typing/backend/db"
)

type User struct {
	Profile     *db.ProfileModel
	Preferences *db.PreferencesModel
}

func NewUser(profile *db.ProfileModel, preferences *db.PreferencesModel) *User {
	return &User{
		Profile:     profile,
		Preferences: preferences,
	}
}
