package db

// PreferencesModel is a model for storing preferences
type PreferencesModel struct {
	ID          string `bson:"_id,omitempty" json:"_id,omitempty"`
	DisplayName string `bson:"displayName" json:"displayName"`
	Email       string `bson:"email" json:"email"`
	Theme       string `bson:"theme" json:"theme"`
}

// NewPreferences is the constructor for PreferencesModel - it creates a new PreferencesModel object
func NewPreferences(id string, displayName string, email string, theme string) *PreferencesModel {

	res := &PreferencesModel{
		ID:          id,
		DisplayName: displayName,
		Email:       email,
		Theme:       theme,
	}
	return res
}
