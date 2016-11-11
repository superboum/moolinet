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

func TestLoadChallengesFromPath(t *testing.T) {
	list, err, warn := LoadChallengesFromPath("../../tests/loadChallengeTest")
	if err != nil || len(list) != 3 || len(warn) != 1 {
		t.Error("Did not load sucessfully", len(list), err, len(warn), warn)
	}
}
