package helpers

import "testing"

var stringLen = 12

func TestGenerateRandomString(t *testing.T) {
	randomString := GenerateRandomString(stringLen)
	if len(randomString) != stringLen {
		t.Error("wrong len string: ", len(randomString))
	}
}
