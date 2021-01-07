package db

import (
	"github.com/hujoseph99/typingBackend/secret"
)

func getMongoURI() string {
	return secret.MongoURI
}
