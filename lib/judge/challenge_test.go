package judge

import (
	"os"
	"testing"
)

func TestNewChallengeFromJSON(t *testing.T) {
	reader, err := os.Open("../../tests/challenge1.json")

	if err != nil {
		t.Error("There was an error while opening the file", err)
		return
	}

	chal, err := NewChallengeFromJSON(reader)
	if err != nil {
		t.Error("There was an error in JSON parsing", err)
		return
	}

	if chal.Title != "Challenge Test #1" || chal.Body != "Hello World" {
		t.Error("There was an error in JSON parsing")
		return
	}

	t.Log(chal)
}
