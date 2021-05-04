package db

import (
	"context"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	newUser := NewUser("username", "password", "email@email.com", "googleID", "facebookID", "githubID", time.Now())
	newPreferences := NewPreferences(newUser.ID, "nickname")
	newProfile := NewProfile(newUser.ID, 18, 19, 20, 21, 22, 23)

	err := RegisterUser(context.Background(), newUser, newProfile, newPreferences)
	if err != nil {
		t.Fatal("Error adding documents")
	}

	if newUser.ID != newPreferences.ID || newUser.ID != newProfile.ID {
		t.Fatal("IDs were not equal")
	}

	foundUser, err := GetUserByID(context.Background(), newUser.ID.Hex(), RegularID)
	if err != nil || foundUser.ID.Hex() != newUser.ID.Hex() ||
		foundUser.Username != newUser.Username || foundUser.Password != newUser.Password ||
		foundUser.Email != newUser.Email {
		t.Fatal("User was not found correctly")
	}

	foundPreferences, err := GetPreferencesByID(context.Background(), newPreferences.ID)
	if err != nil || foundPreferences.ID.Hex() != newPreferences.ID.Hex() ||
		foundPreferences.DisplayName != newPreferences.DisplayName {
		t.Fatal("Preferences was not found correctly")
	}

	foundProfile, err := GetProfileByID(context.Background(), newProfile.ID)
	if err != nil || foundProfile.ID.Hex() != newProfile.ID.Hex() ||
		foundProfile.AverageTPMAllTime != newProfile.AverageTPMAllTime ||
		foundProfile.AverageTPMLast10 != newProfile.AverageTPMLast10 {
		t.Fatal("Profile was not found correctly")
	}

	err = DeleteUser(context.Background(), newUser.ID.Hex(), RegularID)
	if err != nil {
		t.Fatal("The collections were not deleted correctly")
	}
}
