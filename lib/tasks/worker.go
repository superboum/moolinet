package tasks

import "log"

// Worker can executes Jobs in a sequential order.
type Worker struct {
	Jobs *JobQueue
	Stop bool
}

// NewWorker returns a new Worker.
func NewWorker(Jobs *JobQueue) *Worker {
	w := new(Worker)
	w.Jobs = Jobs
	return w
}

// Launch returns immediately, but spawns a new goroutine containing the worker.
func (w *Worker) Launch() {
	go func(w *Worker) {
		for !w.Stop {
			log.Println("Waiting for a new job")
			job := w.Jobs.Get()
			log.Println("Get a new job", job.UUID)
			err := job.Process()
			if err != nil {
				log.Println("An error occured while processing the job: " + err.Error())
			}
			err = job.Callback(job)
			if err != nil {
				log.Println("An error occured in the callback of the job: " + err.Error())
			}
		}
	}(w)
}
