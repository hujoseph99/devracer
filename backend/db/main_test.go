package db

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDatabase()

	os.Exit(m.Run())
}
