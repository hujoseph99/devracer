package secret

import (
	"log"
	"os"
)

var SecretStateString string
var GithubClientID string
var GithubClientSecret string
var MongoURI string

func init() {
	var check bool
	SecretStateString, check = os.LookupEnv("SECRET_STATE_STRING")
	if !check {
		log.Fatal("No environment variables")
	}

	GithubClientID, check = os.LookupEnv("GITHUB_CLIENT_ID")
	if !check {
		log.Fatal("No environment variables")
	}

	GithubClientSecret, check = os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !check {
		log.Fatal("No environment variables")
	}

	MongoURI, check = os.LookupEnv("MONGO_URI")
	if !check {
		log.Fatal("No environment variables")
	}
}
