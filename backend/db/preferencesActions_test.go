package db

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *Client

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
		checkPref, err := client.GetPreferences(context.TODO(), userid)
		if err != nil {
			t.Error("Could not do GetPreferences " + err.Error())
		}

		if !pref.Equals(checkPref) {
			t.Error("GetPreferences does not return the same values")
		}
	}

	client.AddPreferences(context.TODO(), pref)
	getAndCheckPreferences()

	pref.Theme = "dark"
	err = client.UpdatePreferences(context.TODO(), userid, pref)
	if err != nil {
		t.Error("Could not do UpdatePreferences" + err.Error())
	}
	getAndCheckPreferences()

	client.DeletePreferencesByID(context.TODO(), userid)
}

func TestMain(m *testing.M) {
	var err error
	client, err = ConnectToDB(context.TODO())

	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
