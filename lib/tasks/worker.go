package tasks

import "log"

type Worker struct {
	Jobs *JobQueue
	Stop bool
}

func NewWorker(Jobs *JobQueue) *Worker {
	w := new(Worker)
	w.Jobs = Jobs
	return w
}

func (w *Worker) Launch() {
	go func(w *Worker) {
		for w.Stop == false {
			log.Println("Waiting for a new job")
			job := w.Jobs.Get()
			log.Println("Get a new job", job.UUID)
			job.Process()
		}
	}(w)
}
