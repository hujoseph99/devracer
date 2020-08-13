package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProfileModel is a model for user profile data
type ProfileModel struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	TotalWordsTyped   int                `bson:"totalWordsTyped" json:"totalWordsTyped"`
	RacesCompleted    int                `bson:"racesCompleted" json:"racesCompleted"`
	RacesWon          int                `bson:"racesWon" json:"racesWon"`
	MaxTPM            float64            `bson:"maxTPM" json:"maxTPM"`
	AverageTPMAllTime float64            `bson:"averageTPMAllTime" json:"averageTPMAllTime"`
	AverageTPMLast10  float64            `bson:"averageTPMLast10" json:"averageTPMLast10"`
}

// NewProfile is a constructor for the ProfileModel
func NewProfile(id primitive.ObjectID, totalWordsTyped, racesCompleted, racesWon int,
	maxTPM, averageTPMAllTime, averageTPMLast10 float64) *ProfileModel {

	res := &ProfileModel{
		ID:                id,
		TotalWordsTyped:   totalWordsTyped,
		RacesCompleted:    racesCompleted,
		RacesWon:          racesWon,
		MaxTPM:            maxTPM,
		AverageTPMAllTime: averageTPMAllTime,
		AverageTPMLast10:  averageTPMLast10,
	}

	return res
}
