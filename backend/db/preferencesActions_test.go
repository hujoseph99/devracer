package db

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func TestBasic(t *testing.T) {
// 	user := NewUser("a", "pass", "salt", "", "", "", time.Now())
// 	_, err := client.AddUser(context.TODO(), user)
// 	if err != nil {
// 		t.Error("Could not do AddUser" + err.Error())
// 	}
// 	pref := NewPreferences(user.ID, "", "", "")
// 	err = client.AddPreferences(context.TODO(), pref)
// 	if err != nil {
// 		t.Error("Could not do GetPreferences" + err.Error())
// 	}

// 	client.DeleteUserByID(context.TODO(), user.ID)
// 	client.DeletePreferencesByID(context.TODO(), pref.ID)
// }

func TestPreferences(t *testing.T) {
	userid, err := primitive.ObjectIDFromHex("111111111111111111111111")
	if err != nil {
		t.Error("Could not do ObjectIDFromHex " + err.Error())
	}

	pref := NewPreferences(userid, "", "", "")

	var getAndCheckPreferences = func() {
		checkPref, err := GetPreferencesByID(context.TODO(), userid)
		if err != nil {
			t.Error("Could not do GetPreferences " + err.Error())
		}

		if !reflect.DeepEqual(pref, checkPref) {
			t.Error("GetPreferences does not return the same values")
		}
	}

	AddPreferences(context.TODO(), pref)
	getAndCheckPreferences()

	pref.Theme = "dark"
	err = UpdatePreferences(context.TODO(), userid, pref)
	if err != nil {
		t.Error("Could not do UpdatePreferences" + err.Error())
	}
	getAndCheckPreferences()

	DeletePreferencesByID(context.TODO(), userid)
}
