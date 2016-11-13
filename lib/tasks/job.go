package tasks

import (
	"github.com/superboum/moolinet/lib/sandbox"
	"github.com/superboum/moolinet/lib/tools"
)

const (
	READY         = iota
	PROVISIONNING = iota
	IN_PROGRESS   = iota
	DONE          = iota
	FAILED        = iota
)

type Job struct {
	UUID       string
	Image      string
	Executions []Execution
	Status     int
	Progress   chan Execution `json:"-"`
}

func NewJob(image string, template JobTemplate, variables map[string]string) (*Job, error) {
	j := new(Job)

	uuid, err := tools.NewUUID()
	if err != nil {
		return nil, err
	}

	j.UUID = uuid
	j.Image = image
	j.Executions = template.GenerateExecution(variables)
	j.Progress = make(chan Execution, 100)

	return j, nil
}

func (j *Job) Process() error {
	// Run the task
	sandbox, err := sandbox.NewDockerSandbox(j.Image)
	if err != nil {
		return err
	}
	defer sandbox.Destroy()

	for index, exec := range j.Executions {
		out, err := sandbox.Run(exec.Command, exec.Timeout, exec.Network)
		j.Executions[index].Output = out
		j.Executions[index].Error = err

		j.Progress <- j.Executions[index]
		if err != nil {
			break
		}
	}
	close(j.Progress)

	return nil
}
