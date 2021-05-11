package utils_test

import (
	"testing"

	"github.com/hujoseph99/typing/backend/common/utils"
)

func TestFindFirstDifference(t *testing.T) {
	s1 := "The fox jumped over the lazy dog."
	s2 := "The fox jumped"
	s3 := ""
	s4 := s1

	difference := utils.FindFirstDifference(s1, s2)
	if difference != 14 {
		t.Fatal("incorrect difference 1")
	}

	difference = utils.FindFirstDifference(s1, s3)
	if difference != 0 {
		t.Fatal("incorrect difference 2")
	}

	difference = utils.FindFirstDifference(s1, s4)
	if difference != len(s1) {
		t.Fatal("incorrect difference 3")
	}
}
