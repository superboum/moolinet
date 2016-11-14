package tasks

import (
	"strings"
	"testing"
)

func TestNewJob(t *testing.T) {
	jt := JobTemplate{[]Execution{Execution{Command: []string{"cat", "[PATH]"}, Network: true, Timeout: 120}}}
	vars := map[string]string{"[PATH]": "/etc/hosts"}
	j, err := NewJob("superboum/moolinet-golang", jt, vars)
	if err != nil {
		t.Error("Unable to create the job", err)
		return
	}
	err = j.Process()
	if err != nil {
		t.Error("An error occured", err)
		return
	}

	if j.Executions[0].Error != "" {
		t.Error("This error should not occur", j.Executions[0].Error)
	}

	if !strings.Contains(j.Executions[0].Output, "127.0.0.1") {
		t.Error("Output is incorrect ", j.Executions[0].Output)
	}
}
