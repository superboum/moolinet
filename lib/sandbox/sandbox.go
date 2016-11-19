// Package sandbox contains code related to the worker of moolinet.
// It will compile, analyze and execute the code in a sandbox.
package sandbox

import "time"

// Config is the structure used as a parameter for sandboxed commands.
type Config struct {
	Timeout time.Duration
	Network bool
}

// Sandbox is the interface which enable you to run your program independently of the implementation of the sandbox.
type Sandbox interface {
	// Destroy the sandbox on the main system.
	// Usually just after its creation with defer.
	Destroy()

	// Run a command in the sandbox and get its outputs
	// (output, error)
	Run(command []string, config Config) (string, error)

	// Get logs linked to the Sandbox
	// Useful for debugging
	GetLogs() string
}
