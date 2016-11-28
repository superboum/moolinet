package tasks

import (
	"strings"
	"testing"

	"github.com/superboum/moolinet/lib/sandbox"
)

func TestWorker(t *testing.T) {
	// Worker + Queue
	jq := NewJobQueue()
	w := NewWorker(jq)
	w.Launch()

	// Job template + creation
	jt := JobTemplate{[]*Execution{&Execution{
		Command: []string{"cat", "[PATH]"},
		Network: true,
		Timeout: 120}}}
	vars := map[string]string{
		"[PATH]": "/etc/hosts"}
	j, err := NewJob(sandbox.DockerSandboxConfig{Image: "superboum/moolinet-golang"}, jt, vars, func(_ *Job) error { return nil })
	if err != nil {
		t.Error("Couldn't create a job", err)
		return
	}

	// Add the job to the queue
	jq.Add(j)

	for progress := range j.Progress {
		t.Log(progress.Output)
	}

	if j.Executions[0].Error != "" {
		t.Error("Should not be errored", j.Executions[0].Error)
	}
	if !strings.Contains(j.Executions[0].Output, "127.0.0.1") {
		t.Error("Wrong output", j.Executions[0].Output)
	}
}
