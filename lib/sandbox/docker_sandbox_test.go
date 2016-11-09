package sandbox

import (
	"testing"
)

func TestDockerSandbox(t *testing.T) {
	s, err := NewDockerSandbox("superboum/moolinet-golang")

	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	defer s.Destroy()

	output, err := s.Run([]string{"go", "get", "-v", "-d", "github.com/superboum/atuin/..."}, 120, true)
	t.Log("output: " + output)
	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	output, err = s.Run([]string{"go", "get", "-v", "-d", "github.com/superboum/moolinet/..."}, 120, false)
	t.Log("output: " + output)
	if err == nil {
		t.Error("Should throw an error")
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	output, err = s.Run([]string{"go", "install", "-v", "github.com/superboum/atuin/..."}, 120, false)
	t.Log("output: " + output)
	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	output, err = s.Run([]string{"atuin-front"}, 30, false)
	t.Log("output: " + output)
	if err == nil {
		t.Error("Should throw a timeout error")
		return
	}
}
