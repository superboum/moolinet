package judge

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/superboum/moolinet/lib/persistence"
	"github.com/superboum/moolinet/lib/tasks"
	"github.com/superboum/moolinet/lib/tools"
)

// Judge is the entrypoint of this package.
// It stores every component of the judging system (workers, challenges...).
type Judge struct {
	Queue            *tasks.JobQueue
	ActiveJobs       map[string]*tasks.Job
	Challenges       map[string]*Challenge
	PublicChallenges []*Challenge
	Config           *tools.Config
	Warnings         []error
	mutex            sync.Mutex
}

// NewSimpleJudge returns a Judge from configuration.
func NewSimpleJudge(conf *tools.Config) (*Judge, error) {
	j := new(Judge)
	j.Queue = tasks.NewJobQueue()
	j.Config = conf
	j.Warnings = make([]error, 0)
	j.ActiveJobs = make(map[string]*tasks.Job)
	j.mutex = sync.Mutex{}

	err := j.ReloadChallenge()
	if err != nil {
		return nil, err
	}

	if conf.Workers <= 0 {
		conf.Workers = 1
	}

	for i := 0; i < conf.Workers; i++ {
		tasks.NewWorker(j.Queue).Launch()
	}

	return j, nil
}

// ReloadChallenge reloads challenges from disk.
func (j *Judge) ReloadChallenge() error {
	// Load challenges
	chal, warn, err := LoadChallengesFromPath(j.Config.ChallengesPath)
	if err != nil {
		return err
	}
	j.Warnings = append(j.Warnings, warn...)

	j.Challenges = chal
	j.PublicChallenges = GeneratePublicChallenges(chal)

	return nil

}

// Submit submits the code to the queue, returning the waiting job.
func (j *Judge) Submit(slug string, vars map[string]string, u *persistence.User) (*tasks.Job, error) {
	chal, ok := j.Challenges[slug]
	if !ok {
		return nil, errors.New("This challenge does not exist")
	}

	cb := func(currentJob *tasks.Job) error {
		// Job saving
		// @FIXME Challenge and Judge should be in a different package as Judge depends on persistence but persistence should depends on Challenge
		_, err := persistence.NewTerminatedJobFromJob(slug, u, currentJob)

		// Remove job from the list
		// @FIXME There is probably a better solution
		go func() {
			time.Sleep(1 * time.Minute)
			log.Println("Clear job", currentJob.UUID)
			j.mutex.Lock()
			delete(j.ActiveJobs, currentJob.UUID)
			j.mutex.Unlock()
		}()

		return err
	}

	job, err := tasks.NewJob(chal.Docker, chal.Template, vars, cb)
	if err != nil {
		return nil, err
	}

	j.mutex.Lock()
	j.ActiveJobs[job.UUID] = job
	j.mutex.Unlock()
	j.Queue.Add(job)

	return job, nil
}

// GetJob returns the Job with the provided UUID.
// It returns an error is the asked Job is no longer Active.
func (j *Judge) GetJob(UUID string) (*tasks.Job, error) {
	job, ok := j.ActiveJobs[UUID]
	if !ok {
		return nil, errors.New("This job does not exist")
	}

	return job, nil
}
