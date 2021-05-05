package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RefreshToken string             `bson:"refreshToken" json:"refreshToken"`
	UserID       primitive.ObjectID `bson:"userid" json:"userid"`
	ExpiryDate   time.Time          `bson:"expiryDate" json:"expiryDate"`
}

func NewSession(refreshToken string, userID primitive.ObjectID, expiryDate time.Time) *SessionModel {
	res := &SessionModel{
		RefreshToken: refreshToken,
		UserID:       userID,
		ExpiryDate:   expiryDate,
	}

	return res
}
