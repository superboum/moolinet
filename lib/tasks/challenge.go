package tasks

import (
	"encoding/json"
	"io"
)

type Challenge struct {
	Title      string
	Body       string
	Image      string
	Executions []Execution
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
