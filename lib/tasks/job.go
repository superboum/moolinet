package tasks

import (
	"log"
)

type Job struct {
	UID        string
	Image      string
	Executions []Execution
	Status     int
}

func NewJob(image string, et ExecutionsTemplate) {
	j := new(Job)
	j.Image = image
	j.Executions = et.GenerateExecution()
	return et
}

func (j *Job) Process() {
	// Run the task
	sandbox, err := sandbox.NewDockerSandbox(job.Image)
	if err != nil {
		Log.Println("Unable to create the sandbox")
	}

	for _, exec := range sandbox.Executions {
		exec.Output, exec.Error = sandbox.Run(exec.Command, exec.Network, exec.Timeout)
	}
}
