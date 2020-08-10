package db

import (
	"time"
)

// UserModel is a model for user data that will be used for authentication
type UserModel struct {
	ID           string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Username     string    `bson:"username" json:"username"`
	Password     string    `bson:"password" json:"password"`
	Salt         string    `bson:"salt" json:"salt"`
	GoogleID     string    `bson:"googleID,omitempty" json:"googleID,omitempty"`
	GithubID     string    `bson:"githubID,omitempty" json:"githubID,omitempty"`
	FacebookID   string    `bson:"facebookID,omitempty" json:"facebookID,omitempty"`
	RegisterDate time.Time `bson:"registerDate" json:"registerDate"`
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
