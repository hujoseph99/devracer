package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// emulated enum to represent id types
const (
	RegularID  = iota
	GoogleID   = iota
	GithubID   = iota
	FacebookID = iota
)

// UserModel is a model for user data that will be used for authentication
type UserModel struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Salt         string             `bson:"salt" json:"salt"`
	GoogleID     string             `bson:"googleID" json:"googleID"`
	GithubID     string             `bson:"githubID" json:"githubID"`
	FacebookID   string             `bson:"facebookID" json:"facebookID"`
	RegisterDate time.Time          `bson:"registerDate" json:"registerDate"`
}

// NewUser is the constructor for User - it creates a new User object
func NewUser(username, password, salt, googleID, githubID, facebookID string,
	registerDate time.Time) *UserModel {

	res := &UserModel{
		Username:     username,
		Password:     password,
		Salt:         salt,
		GoogleID:     googleID,
		GithubID:     githubID,
		FacebookID:   facebookID,
		RegisterDate: registerDate,
	}

	return res
}
