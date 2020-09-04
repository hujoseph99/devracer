package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

const saltLength = 16

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// Gotten from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// encryptString will encrypt a given string using an AES cipher.
//	This is an internal method that will typically be used for
//	generating a state string for OAuth applications
func encryptString(plaintext string) (string, error) {
	keyBytes := []byte(getSecretStateString())
	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	res := gcm.Seal(nonce, nonce, plaintextBytes, nil)
	return string(res), nil
}

// decryptString will decrypt a given string assuming that an
//	AES cipher was used (in particular the one above). This
// is an internal method that will typically be used for decrypting
//	state strings for OAuth applications
func decryptString(ciphertext string) (string, error) {
	keyBytes := []byte(getSecretStateString())

	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("Invalid ciphertext")
	}

	ciphertextBytes := []byte(ciphertext)
	nonce, ciphertextBytes := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext), nil
}

// generateEcryptedOAuthStateString will geneate a state string
//	that can be passed along OAuth requests
func generateEncryptedOAuthStateString() (string, error) {
	salt, err := generateRandomString(saltLength)
	if err != nil {
		return "", err
	}

	stateString := getSecretStateString()
	tempStateString := stateString + salt
	encryptedState, err := encryptString(tempStateString)
	if err != nil {
		return "", err
	}
	return encryptedState, nil
}

func verifyOAuthStateString(state string) (bool, error) {
	plaintext, err := decryptString(state)
	if err != nil {
		return false, err
	}

	stateString := getSecretStateString()
	return plaintext[:len(plaintext)-saltLength] == stateString, nil
}
