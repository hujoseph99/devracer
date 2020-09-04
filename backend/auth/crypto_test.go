package auth

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	testnum := 10
	stringLength := 15
	arr := make([]string, testnum)
	for i := range arr {
		rand, err := generateRandomString(stringLength)
		if err != nil {
			t.Error("Error generating a random string")
		}
		arr[i] = rand
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				t.Error("The two random strings are the same")
			}
		}
	}
}

func TestEncryptString(t *testing.T) {
	testString := "test-string"
	res, err := encryptString(testString)
	if err != nil {
		t.Error("Error encrypting string")
	}
	if res == testString {
		t.Error("String was not encrypted")
	}
}

func TestDecryptString(t *testing.T) {
	testString := "test-string"
	res, err := encryptString(testString)
	if err != nil {
		t.Error("Error encrypting string")
	}

	decrypted, err := decryptString(res)
	if err != nil {
		t.Error("Error decrypting string")
	}
	if decrypted != testString {
		t.Errorf("plaintext string (%v) and decrypted string (%v) do not match",
			testString, decrypted)
	}
}

func TestOAuthWorkflow(t *testing.T) {
	state, err := generateEncryptedOAuthStateString()
	if err != nil {
		t.Error("Error generating OAuth state string")
	}

	res, err := verifyOAuthStateString(state)
	if err != nil || !res {
		t.Error("Error verifying the OAuth state string")
	}
}
