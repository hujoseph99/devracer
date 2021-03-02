package db

import (
	"github.com/hujoseph99/typing/backend/secret"
)

func getMongoURI() string {
	return secret.MongoURI
}
