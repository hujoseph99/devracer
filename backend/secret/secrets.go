package secret

import (
	"log"
	"os"
	"strconv"
)

var SecretStateString string
var GithubClientID string
var GithubClientSecret string
var MongoURI string
var FrontendCallback string
var Production bool
var FrontendHostname string

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

	FrontendCallback, check = os.LookupEnv("FRONTEND_CALLBACK")
	if !check {
		log.Fatal("No environment variables")
	}

	FrontendHostname, check = os.LookupEnv("FRONTEND_HOSTNAME")
	if !check {
		log.Fatal("No environment variables")
	}

	prodtest, check := os.LookupEnv("PRODUCTION")
	if !check {
		log.Fatal("No environment variables")
	}

	var err error
	Production, err = strconv.ParseBool(prodtest)
	if err != nil {
		log.Fatal("invalid PRODUCTION value")
	}
}
