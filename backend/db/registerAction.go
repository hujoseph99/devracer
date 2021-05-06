package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser will register a user given User, Profile, and Preferences models. If it's
// unsuccessful, then nothing will be added and an error will be returned. It will also
// add the correct ID to the profile model and preferences model (the same one as the user model)
func RegisterUser(ctx context.Context, user *UserModel, profile *ProfileModel,
	preferences *PreferencesModel) error {

	if user.ID != profile.ID || user.ID != preferences.ID {
		return fmt.Errorf("the given ids were not consistent")
	}

	session, err := db.StartSession()
	if err != nil {
		return err
	}

	if err = session.StartTransaction(); err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = AddUser(sc, user)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		// add id to the objects because only know the id after mongo inserts the user
		profile.ID = user.ID
		preferences.ID = user.ID

		err = AddProfile(sc, profile)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		err = AddPreferences(sc, preferences)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser will delete a user given its ID. Along with the user, it will also delete their
// corresponing profile and preferences from the db. It will return an error if there is one.
// If it's successful, nothing will be returned.
func DeleteUser(ctx context.Context, userID string, idType int) error {
	// first get the user by ID.
	user, err := GetUserByID(ctx, userID, idType)
	if err != nil {
		return err
	}

	session, err := db.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		// delete the user by its regularID always - need it anyway for the others anyway
		err = DeleteUserByID(sc, user.ID.Hex(), RegularID)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		err = DeletePreferencesByID(sc, user.ID)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		err = DeleteProfileByID(sc, user.ID)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
