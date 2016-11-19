package sandbox

import (
	"testing"
	"time"
)

func TestDockerSandbox(t *testing.T) {
	s, err := NewDockerSandbox("superboum/moolinet-golang")

	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	defer s.Destroy()

	c := Config{
		Timeout: 2 * time.Minute,
		Network: true,
	}
	output, err := s.Run([]string{"go", "get", "-v", "-d", "github.com/superboum/atuin/..."}, c)
	t.Log("output: " + output)
	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	c.Network = false
	output, err = s.Run([]string{"go", "get", "-v", "-d", "github.com/superboum/moolinet/..."}, c)
	t.Log("output: " + output)
	if err == nil {
		t.Error("Should throw an error")
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	output, err = s.Run([]string{"go", "install", "-v", "github.com/superboum/atuin/..."}, c)
	t.Log("output: " + output)
	if err != nil {
		t.Error("Unexpected error", err)
		return
	}
	if output == "" {
		t.Error("Output should not be null")
		return
	}

	c.Timeout = 30 * time.Second
	output, err = s.Run([]string{"atuin-front"}, c)
	t.Log("output: " + output)
	if err == nil {
		t.Error("Should throw a timeout error")
		return
	}
}
