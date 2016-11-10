package tasks

type Worker struct {
	Jobs JobQueue
}

func NewWorker(Jobs JobQueue) *Worker {
	w := new(Worker)
	w.Jobs = Jobs
	return w
}

func (w *Worker) Launch() {
	for {
		job := w.Jobs.Get()
		job.Process()
	}
}
