package api

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserReturnToClient is a struct that will be used to return data back to the client.
// It will exclude sensitive data such as emails and passwords that the db typically
// returns.
type UserReturnToClient struct {
	Username     string    `bson:"username" json:"username"`
	Nickname     string    `bson:"nickname" json:"nickname"`
	Wpm          int       `bson:"wpm,minsize" json:"wpm"`
	RegisterDate time.Time `bson:"register_date" json:"register_date"`
}

// NewUserReturnToClient will create a new user struct that is meant to be returned
// to clients
// func NewUserReturnToClient(user *db.UserModel) (*UserReturnToClient, error) {
// 	if user == nil {
// 		return nil, fmt.Errorf("The given user is invalid")
// 	}

// 	res := &UserReturnToClient{
// 		Username:     user.Username,
// 		Nickname:     user.Nickname,
// 		Wpm:          user.Wpm,
// 		RegisterDate: user.RegisterDate,
// 	}

// 	return res, nil
// }

// JWTExpireTime will be the amount of time before the JWT expires
const JWTExpireTime = time.Minute * 15

// TODO: Change secret and put into environment file
const tempSecret = "abc123abc123"

// convertToMapClaims will take a User object that is meant to be returned
// to the client and conver it to a jwt.MapClaims to be used for jwt
// encoding
func (user *UserReturnToClient) convertToMapClaims() *jwt.MapClaims {
	fields := reflect.TypeOf(*user)
	values := reflect.ValueOf(*user)

	atClaims := jwt.MapClaims{}

	// get tag names (from json) and values then add to atClaims
	for i := 0; i < values.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		tag, ok := field.Tag.Lookup("json")
		if ok {
			atClaims[tag] = value.Interface()
		}
	}

	return &atClaims
}

// func parseToken(tokenString string) {
// 	hmacSecret := []byte(tempSecret)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return hmacSecret, nil
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println(claims)
// 	}

// }

// convertToJwt will convert a User object that is meant to be returned
// to the client and convert it to a jwt
func (user *UserReturnToClient) convertToJwt() (string, error) {
	atClaims := user.convertToMapClaims()
	(*atClaims)["exp"] = time.Now().Add(JWTExpireTime).Unix() // Set expire time
	res := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tempSecretConverted := []byte(tempSecret)
	token, err := res.SignedString(tempSecretConverted)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parseToken(token)
	return token, nil
}
