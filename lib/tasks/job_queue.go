package tasks

// JobQueue is a generic structure for job management (FIFO).
type JobQueue struct {
	queue chan *Job
}

// NewJobQueue returns a clean JobQueue.
func NewJobQueue() *JobQueue {
	jq := new(JobQueue)
	jq.queue = make(chan *Job, 200)

	return jq
}

// Add adds the provided job to the queue.
func (jq *JobQueue) Add(job *Job) {
	jq.queue <- job
}

// Get returns one ready-to-go job.
func (jq *JobQueue) Get() *Job {
	job := <-jq.queue
	return job
}
