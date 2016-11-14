package judge

import (
	"errors"

	"github.com/superboum/moolinet/lib/tasks"
	"github.com/superboum/moolinet/lib/tools"
)

type Judge struct {
	Queue            *tasks.JobQueue
	Worker           *tasks.Worker
	ActiveJobs       map[string]*tasks.Job
	Challenges       map[string]*Challenge
	PublicChallenges []*Challenge
	Config           *tools.Config
	Warnings         []error
}

func NewSimpleJudge(conf *tools.Config) (*Judge, error) {
	j := new(Judge)
	j.Queue = tasks.NewJobQueue()
	j.Worker = tasks.NewWorker(j.Queue)
	j.Config = conf
	j.Warnings = make([]error, 0)
	j.ActiveJobs = make(map[string]*tasks.Job)

	err := j.ReloadChallenge()
	if err != nil {
		return nil, err
	}

	j.Worker.Launch()
	return j, nil
}

func (j *Judge) ReloadChallenge() error {
	// Load challenges
	chal, err, warn := LoadChallengesFromPath(j.Config.ChallengesPath)
	if err != nil {
		return err
	}
	j.Warnings = append(j.Warnings, warn...)

	j.Challenges = chal
	j.PublicChallenges = GeneratePublicChallenges(chal)

	return nil

}

func (j *Judge) Submit(slug string, vars map[string]string) (*tasks.Job, error) {
	chal, ok := j.Challenges[slug]
	if !ok {
		return nil, errors.New("This challenge does not exist")
	}
	job, err := tasks.NewJob(chal.Image, chal.Template, vars)
	if err != nil {
		return nil, err
	}
	j.Queue.Add(job)
	j.ActiveJobs[job.UUID] = job

	return job, nil
}

func (j *Judge) GetJob(UUID string) (*tasks.Job, error) {
	job, ok := j.ActiveJobs[UUID]
	if !ok {
		return nil, errors.New("This job does not exist")
	}

	return job, nil
}
