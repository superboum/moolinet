package judge

import (
	"github.com/superboum/moolinet/lib/tools"
	"testing"
)

func TestJudge(t *testing.T) {
	judge, err := NewSimpleJudge(&tools.Config{ChallengesPath: "../../tests/loadChallengeTest"})
	if err != nil {
		t.Error("Unable to create judge", err)
		return
	}
	job, err := judge.Submit("challenge2ok", map[string]string{"[GIT-REPO]": "https://github.com/superboum/atuin"})
	if err != nil {
		t.Error("Couldn't create job", err)
		return
	}

	for progress := range job.Progress {
		if progress.Error != nil {
			t.Error("Command should not fail", progress.Error, progress.Output)
			return
		}
	}
}
