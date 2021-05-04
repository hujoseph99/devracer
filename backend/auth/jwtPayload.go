package auth

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hujoseph99/typing/backend/secret"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type jwtPayload struct {
	UserId string `json:"userid"`
}

func newJwtPayload(id primitive.ObjectID) *jwtPayload {
	res := &jwtPayload{
		UserId: id.Hex(),
	}
	return res
}

// JWTExpireTime will be the amount of time before the JWT expires
const JWTExpireTime = time.Minute * 15

// convertToMapClaims will take a User object that is meant to be returned
// to the client and conver it to a jwt.MapClaims to be used for jwt
// encoding
func (user *jwtPayload) convertToMapClaims() *jwt.MapClaims {
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

func (user *jwtPayload) convertToJwt() (string, error) {
	atClaims := user.convertToMapClaims()
	(*atClaims)["exp"] = time.Now().Add(JWTExpireTime).Unix() // Set expire time
	res := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tempSecretConverted := []byte(secret.SecretStateString)
	token, err := res.SignedString(tempSecretConverted)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parseToken(token)
	return token, nil
}

func parseToken(tokenString string) {
	hmacSecret := []byte(secret.SecretStateString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}

}
