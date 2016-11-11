package judge

import (
	"encoding/json"
	"io"

	"github.com/superboum/moolinet/lib/tasks"
)

type Challenge struct {
	Title    string
	Body     string
	Image    string
	Template tasks.JobTemplate
}

func NewChallengeFromJSON(reader io.Reader) (*Challenge, error) {
	decoder := json.NewDecoder(reader)

	chal := new(Challenge)
	err := decoder.Decode(&chal)
	if err != nil {
		return nil, err
	}

	return chal, nil
}
