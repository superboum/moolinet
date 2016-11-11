package judge

import (
	"github.com/superboum/moolinet/lib/tasks"
	"github.com/superboum/moolinet/lib/tools"
)

type Judge struct {
	Queue      *tasks.JobQueue
	Worker     *tasks.Worker
	Challenges map[string]Challenge
	Config     *tools.Config
}

func NewSimpleJudge(conf *tools.Config) *Judge {
	j := new(Judge)
	j.Queue = tasks.NewJobQueue()
	j.Worker = tasks.NewWorker(j.Queue)
	j.Config = conf

	j.ReloadChallenge()
	j.Worker.Launch()

	return j
}

func (j *Judge) ReloadChallenge() {
	// Load challenges
	// j.Challenges = LoadChallengesFromPath(j.Config.ChallengesPath)
}

func (j *Judge) Submit(challengeID string, vars []string) *tasks.Job {
	return nil
}
