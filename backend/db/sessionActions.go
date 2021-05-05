package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSessionBySessionString will get a session from the db by a given session string.
func GetSessionBySessionString(ctx context.Context, sessionString string) (*SessionModel, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSessions)

	filter := bson.M{"refreshToken": sessionString}

	var session SessionModel
	err := collection.FindOne(ctx, filter).Decode(&session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// GetPreferencesByID returns the session for a given id.
func GetSessionByID(ctx context.Context, id primitive.ObjectID) (*SessionModel, error) {
	collection := db.Database(DatabaseTypers).Collection(CollectionsSessions)

	var session SessionModel
	err := getDocumentFromCollectionByID(ctx, collection, id, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func UpdateSession(ctx context.Context, id primitive.ObjectID, session *SessionModel) error {

}

// // UpdatePreferences updates the preferences for a user.
// func UpdatePreferences(ctx context.Context, id primitive.ObjectID, pref *PreferencesModel) error {
// 	collection := db.Database(DatabaseTypers).Collection(CollectionsPreferences)

// 	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": pref})

// 	if result.Err() != nil {
// 		return result.Err()
// 	}
// 	return nil
// }

// // AddPreferences adds a given pref to a mongo client.  Will return the error if there is an error.
// func AddPreferences(ctx context.Context, pref *PreferencesModel) error {
// 	collection := db.Database(DatabaseTypers).Collection(CollectionsPreferences)

// 	_, err := addDocumentToCollection(ctx, collection, pref)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // DeletePreferencesByID will delete a preferences by the given id.  If it is successful, then
// // it will return a nil error, otherwise it will return an error.
// func DeletePreferencesByID(ctx context.Context, id primitive.ObjectID) error {
// 	collection := db.Database(DatabaseTypers).Collection(CollectionsPreferences)

// 	err := deleteFromCollectionByID(ctx, collection, id)
// 	return err
// }
