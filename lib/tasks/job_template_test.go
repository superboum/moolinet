package tasks

import (
	"testing"
)

func TestGenerateExecution(t *testing.T) {
	jt := JobTemplate{[]Execution{Execution{Command: []string{"ls", "[FOLDER]"}}}}
	exec := jt.GenerateExecution(map[string]string{"[FOLDER]": "hello", "[TEST]": "hallo"})
	if exec[0].Command[1] != "hello" {
		t.Error("Templating failed: expecting hello but get", exec[0].Command[1])
	}
	if jt.Executions[0].Command[1] != "[FOLDER]" {
		t.Error("Templating modified the template: expecting [FOLDER] but get", jt.Executions[0].Command[1])
	}
}
