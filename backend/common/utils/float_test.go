package utils_test

import (
	"testing"

	"github.com/hujoseph99/typing/backend/common/utils"
)

func TestRoundTo(t *testing.T) {
	n := float64(10.3523)
	rounded := utils.RoundFloor(n, 2)
	if rounded != float64(10.35) {
		t.Fatal("rounded incorrectly")
	}

	rounded = utils.RoundFloor(n, 1)
	if rounded != float64(10.3) {
		t.Fatal("rounded incorrectly")
	}
}
