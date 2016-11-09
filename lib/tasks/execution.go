package tasks

import (
	"encoding/json"
	"io"
)

const (
	WAIT   = iota
	LAUNCH = iota
	DONE   = iota
)

type Execution struct {
	Command []string
	Network bool
	Timeout int
	Status  int
	Output  string
	Error   error
}

func NewExecutionFromJSON(reader io.Reader) (*Execution, error) {
	decoder := json.NewDecoder(reader)

	exec := new(Execution)
	for decoder.More() {
		err := decoder.Decode(&exec)
		if err != nil {
			return nil, err
		}
	}

	return exec, nil
}
