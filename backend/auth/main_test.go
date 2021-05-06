package auth

import (
	"os"
	"testing"

	"github.com/hujoseph99/typing/backend/db"
)

func TestMain(m *testing.M) {
	db.InitDatabase()

	os.Exit(m.Run())
}
