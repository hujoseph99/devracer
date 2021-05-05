package db

import (
	"time"

	"github.com/dchest/uniuri"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DefaultExpiryTime = time.Hour * 24
const RememberMeExpiryTime = time.Hour * 24 * 7

type SessionModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RefreshToken string             `bson:"refreshToken" json:"refreshToken"`
	UserID       primitive.ObjectID `bson:"userid" json:"userid"`
	ExpiryDate   time.Time          `bson:"expiryDate" json:"expiryDate"`
	Remember     bool               `bson:"remember" json:"remember"`
}

func NewSession(refreshToken string, userID primitive.ObjectID, expiryDate time.Time, remember bool) *SessionModel {
	res := &SessionModel{
		RefreshToken: refreshToken,
		UserID:       userID,
		ExpiryDate:   expiryDate.UTC().Round(time.Millisecond), // mongodb converts to UTC, make it consistent with our models
		Remember:     remember,
	}

	return res
}

// Generates a cryptographically safe refresh token
func GenerateRefreshToken() string {
	return uniuri.NewLen(uniuri.UUIDLen) // length of 20 characters
}

// Get new expiry gets a new expiry time for us in the corect format. It will add onto the current
// date depending on if the user has the remember me option enabled.
func GetNewExpiryTime(remember bool) time.Time {
	if remember {
		return time.Now().UTC().Round(time.Millisecond).Add(RememberMeExpiryTime)
	}
	return time.Now().UTC().Round(time.Millisecond).Add(DefaultExpiryTime)
}
