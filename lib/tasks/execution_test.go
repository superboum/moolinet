package tasks

import (
	"os"
	"testing"
)

func TestNewExecutionFromJSON(t *testing.T) {
	reader, err := os.Open("../../tests/execution1.json")
	if err != nil {
		t.Error("There was an error while opening the file", err)
		return
	}
	exec, _ := NewExecutionFromJSON(reader)
	if exec.Command[0] != "ls" {
		t.Error("There was an error in JSON parsing")
		return
	}
}
