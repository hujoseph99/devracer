package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	slice1, err := generateRandomBytes(15)
	if err != nil {
		t.Error("Could not generate random bytes")
	}

	slice2, err := generateRandomBytes(15)
	if err != nil {
		t.Error("Could not generate random bytes")
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return
		}
	}
	t.Error("The slices were the same")
}

func TestGenerateStateOAuthCookie(t *testing.T) {
	w := httptest.NewRecorder()
	state, err := generateAndSetStateOAuthCookie(w)
	if err != nil || state == "" {
		t.Errorf("Could not create state string for OAuth cookie")
	}
}
