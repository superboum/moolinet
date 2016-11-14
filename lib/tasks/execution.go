package tasks

import (
	"encoding/json"
	"io"
)

type Execution struct {
	Description string
	Command     []string
	Network     bool
	Timeout     int
	Output      string
	Error       string
	Run         bool
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

func (e *Execution) DeepCopy() Execution {
	f := Execution{}

	f.Command = make([]string, len(e.Command))
	copy(f.Command, e.Command)

	f.Description = e.Description
	f.Network = e.Network
	f.Timeout = e.Timeout
	f.Output = e.Output
	f.Error = e.Error
	f.Run = e.Run

	return f
}
