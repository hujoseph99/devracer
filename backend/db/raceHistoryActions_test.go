package db

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestRaceHistory(t *testing.T) {
	raceParticipants := make([]RaceParticipantModel, 0)
	raceParticipants = append(raceParticipants, *NewRaceParticipant("testplayer", 100, 120, 50, 10, 3))
	raceHistory := NewRaceHistory("111111111111111111", raceParticipants, time.Now().UTC().Round(time.Millisecond))

	var getAndCheckRaceHistory = func() {
		checkRaceHistory, err := GetRaceHistoryByID(context.TODO(), raceHistory.ID)
		if err != nil {
			t.Fatal("Could not do GetRaceHistory " + err.Error())
		}

		if !reflect.DeepEqual(raceHistory, checkRaceHistory) {
			t.Fatal("GetRaceHistory does not return the same values")
		}
	}

	AddRaceHistory(context.TODO(), raceHistory)
	getAndCheckRaceHistory()

	raceHistory.SnippetID = "changedsnippet"
	err := UpdateRaceHistory(context.TODO(), raceHistory.ID, raceHistory)
	if err != nil {
		t.Fatal("Could not do UpdateRaceHistory" + err.Error())
	}
	getAndCheckRaceHistory()

	DeleteRaceHistoryByID(context.TODO(), raceHistory.ID)
}
