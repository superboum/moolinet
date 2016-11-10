package tasks

type JobQueue struct {
	Queue chan *Job
}

func NewJobQueue() *JobQueue {
	jq := new(JobQueue)
	jq.Queue = make(chan *Job)

	return jq
}

func (jq *JobQueue) Add(job *Job) {
	jq.Queue <- job
}

func (jq *JobQueue) Get() *Job {
	job := <-jq.Queue
	return job
}
