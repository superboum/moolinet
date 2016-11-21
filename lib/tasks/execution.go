package tasks

import (
	"encoding/json"
	"io"
)

// Execution contains the data related to a particular execution (command) to be
// run in a sandbox.
type Execution struct {
	Description string
	Command     []string
	Expected    string
	Timeout     int
	Output      string
	Error       string
	Run         bool
	Network     bool
	Public      bool // Wether or not the output shall be safely returned to user
}

// NewExecutionFromJSON unmarshals an Execution from JSON input stream.
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

// DeepCopy returns a safe copy of the current Execution.
func (e *Execution) DeepCopy() *Execution {
	f := &Execution{}

	f.Command = make([]string, len(e.Command))
	copy(f.Command, e.Command)

	f.Description = e.Description
	f.Network = e.Network
	f.Timeout = e.Timeout
	f.Output = e.Output
	f.Expected = e.Expected
	f.Error = e.Error
	f.Run = e.Run
	f.Public = e.Public

	return f
}
