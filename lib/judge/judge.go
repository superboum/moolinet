package judge

import (
	"github.com/superboum/moolinet/lib/tasks"
	"github.com/superboum/moolinet/lib/tools"
)

type Judge struct {
	Queue      *tasks.JobQueue
	Worker     *tasks.Worker
	Challenges map[string]*Challenge
	Config     *tools.Config
	Warnings   []error
}

func NewSimpleJudge(conf *tools.Config) (*Judge, error) {
	j := new(Judge)
	j.Queue = tasks.NewJobQueue()
	j.Worker = tasks.NewWorker(j.Queue)
	j.Config = conf
	j.Warnings = make([]error, 0)

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
	return nil

}

func (j *Judge) Submit(slug string, vars map[string]string) (*tasks.Job, error) {
	job, err := tasks.NewJob(j.Challenges[slug].Image, j.Challenges[slug].Template, vars)
	if err != nil {
		return nil, err
	}
	j.Queue.Add(job)

	return job, nil
}