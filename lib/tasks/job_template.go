package tasks

import (
	"strings"
)

// JobTemplate holds generic executions.
type JobTemplate struct {
	Executions []*Execution
}

// GenerateExecution is a basic templating function:
// it replaces tokens (eg: [URL]) by a value (eg: http://example.com)
func (jt *JobTemplate) GenerateExecution(variables map[string]string) []*Execution {
	variables["[_ABOUT]"] = "moolinet"
	templated := make([]*Execution, len(jt.Executions))

	for index, exec := range jt.Executions {
		templated[index] = exec.DeepCopy()
		for index2 := range templated[index].Command {
			for token, value := range variables {
				templated[index].Command[index2] = strings.Replace(templated[index].Command[index2], token, value, -1)
			}
		}
	}
	return templated
}
