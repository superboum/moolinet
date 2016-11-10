package tasks

import (
	"testing"
)

func TestWorker(t *testing.T) {
	/*
		// Worker + Queue
		t.Log("Creation of the queue and the worker")
		jq := NewJobQueue()
		w := NewWorker(jq)
		t.Log("Launch the worker")
		w.Launch()

		// Job template + creation
		t.Log("Creation of the job")
		jt := JobTemplate{[]Execution{Execution{
			Command: []string{"cat", "[PATH]"},
			Network: true,
			Timeout: 120}}}
		vars := map[string]string{
			"[PATH]": "/etc/host"}
		j, err := NewJob("superboum/moolinet-golang", jt, vars)
		if err != nil {
			t.Error("Couldn't create a job", err)
			return
		}

		t.Log("Add the job to the worker")
		// Add the job to the queue
		jq.Add(j)

		t.Log("Wait for the output of the job")
		for progress := range j.Progress {
			t.Log(progress.Output)
		}
	*/
}
