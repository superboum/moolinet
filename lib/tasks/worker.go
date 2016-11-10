package tasks

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
			job := w.Jobs.Get()
			job.Process()
		}
	}(w)
}
