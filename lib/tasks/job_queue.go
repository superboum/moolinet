package tasks

import (
	"../sandbox"
)

type JobQueue struct {
	Queue []Job
}

func NewJobQueue() *JobQueue {
	jq = new(JobQueue)
	jq.Queue = make([]Job, 0)

	return jq
}

func (jq *JobQueue) Add(job Job) {
	jq <- job
}

func (jq *JobQueue) Length() int {
	return len(Queue)
}
