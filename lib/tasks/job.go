package tasks

import (
	"errors"
	"regexp"
	"time"

	"github.com/superboum/moolinet/lib/sandbox"
	"github.com/superboum/moolinet/lib/tools"
)

// ErrUnexpectedOutput is returned when a command does not return the expected result.
var ErrUnexpectedOutput = errors.New("Unexpected output")

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
	Executions []*Execution
	Status     int
	Progress   chan *Execution    `json:"-"`
	Callback   func(j *Job) error `json:"-"`
	Variables  map[string]string  `json:"-"`
}

// NewJob creates a new Job from standard parameters.
func NewJob(config interface{}, template JobTemplate, variables map[string]string, cb func(j *Job) error) (*Job, error) {
	j := new(Job)

	uuid, err := tools.NewUUID()
	if err != nil {
		return nil, err
	}

	j.UUID = uuid
	j.Config = config
	j.Executions = template.GenerateExecution(variables)
	j.Progress = make(chan *Execution, 100)
	j.Status = JobStatusInQueue
	j.Callback = cb
	j.Variables = variables

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
	for _, exec := range j.Executions {
		config := sandbox.Config{
			Timeout: time.Duration(exec.Timeout) * time.Second,
			Network: exec.Network,
		}
		out, err := s.Run(exec.Command, config)
		exec.Output = out
		exec.Run = true

		// Check output content
		r := regexp.MustCompile(exec.Expected)
		if err == nil && !r.MatchString(out) {
			err = ErrUnexpectedOutput
		}

		if err != nil {
			exec.Error = err.Error()
			j.Status = JobStatusFailed
		}

		j.Progress <- exec
	}
	close(j.Progress)

	if j.Status == JobStatusInProgress {
		j.Status = JobStatusSuccess
	}

	return nil
}
