package tasks

import (
	"time"

	"github.com/superboum/moolinet/lib/sandbox"
	"github.com/superboum/moolinet/lib/tools"
)

// These constants represent the different status available for a particular job.
const (
	JobStatusInQueue = iota
	JobStatusProvisionning
	JobStatusInProgress
	JobStatusSuccess
	JobStatusFailed
)

// Job is the structure representing a set of executions to be executed in a sandbox.
type Job struct {
	UUID       string
	Config     interface{}
	Executions []Execution
	Status     int
	Progress   chan Execution `json:"-"`
}

// NewJob creates a new Job from standard parameters.
func NewJob(config interface{}, template JobTemplate, variables map[string]string) (*Job, error) {
	j := new(Job)

	uuid, err := tools.NewUUID()
	if err != nil {
		return nil, err
	}

	j.UUID = uuid
	j.Config = config
	j.Executions = template.GenerateExecution(variables)
	j.Progress = make(chan Execution, 100)
	j.Status = JobStatusInQueue

	return j, nil
}

// Process starts the Job, running every execution.
func (j *Job) Process() error {
	j.Status = JobStatusProvisionning
	s, err := sandbox.NewDockerSandbox(j.Config.(sandbox.DockerSandboxConfig))
	if err != nil {
		j.Status = JobStatusFailed
		return err
	}
	defer s.Destroy()

	j.Status = JobStatusInProgress
	for index, exec := range j.Executions {
		config := sandbox.Config{
			Timeout: time.Duration(exec.Timeout) * time.Second,
			Network: exec.Network,
		}
		out, err := s.Run(exec.Command, config)
		j.Executions[index].Output = out
		j.Executions[index].Run = true
		if err != nil {
			j.Executions[index].Error = err.Error()
		}

		j.Progress <- j.Executions[index]
		if err != nil {
			j.Status = JobStatusFailed
		}
	}
	close(j.Progress)

	if j.Status == JobStatusInProgress {
		j.Status = JobStatusSuccess
	}
	return nil
}
