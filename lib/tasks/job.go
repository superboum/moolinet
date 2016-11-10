package tasks

import (
	"github.com/superboum/moolinet/lib/sandbox"
	"github.com/superboum/moolinet/lib/tools"
	"log"
)

type Job struct {
	UUID       string
	Image      string
	Executions []Execution
	Status     int
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

	return j, nil
}

func (j *Job) Process() {
	// Run the task
	sandbox, err := sandbox.NewDockerSandbox(j.Image)
	if err != nil {
		log.Println("Unable to create the sandbox")
	}

	for _, exec := range j.Executions {
		exec.Output, exec.Error = sandbox.Run(exec.Command, exec.Timeout, exec.Network)
	}
}
