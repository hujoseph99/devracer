package auth

import (
	"github.com/hujoseph99/typingBackend/secret"
)

func getSecretStateString() string {
	return secret.SecretStateString
}

func getGithubClientID() string {
	return secret.GithubClientID
}

func getGithubClientSecret() string {
	return secret.GithubClientSecret
}
