package judge

import (
	"testing"

	"github.com/superboum/moolinet/lib/persistence"
	"github.com/superboum/moolinet/lib/tools"
)

func TestJudge(t *testing.T) {
	judge, err := NewSimpleJudge(&tools.Config{ChallengesPath: "../../tests/loadChallengeTest"})
	if err != nil {
		t.Error("Unable to create judge", err)
		return
	}
	job, err := judge.Submit("02-challenge-ok", map[string]string{"[GIT-REPO]": "https://github.com/superboum/atuin"}, &persistence.User{})
	if err != nil {
		t.Error("Couldn't create job", err)
		return
	}

	for progress := range job.Progress {
		if progress.Error != "" {
			t.Error("Command should not fail", progress.Error, progress.Output)
			return
		}
	}
}

func TestJudgeChallengeNotFound(t *testing.T) {
	judge, err := NewSimpleJudge(&tools.Config{ChallengesPath: "../../tests/loadChallengeTest"})
	if err != nil {
		t.Error("Unable to create judge", err)
		return
	}
	job, err := judge.Submit("challenge2ok", map[string]string{"[GIT-REPO]": "https://github.com/superboum/atuin"}, &persistence.User{})
	if err.Error() != "This challenge does not exist" || job != nil {
		t.Error("The only returned error should be about a challenge which was not found. Get: ", err, job)
		return
	}

}
