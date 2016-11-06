// Libraries related to the worker of moolinet
// It will compile, analyse and execute the code in a sandbox
package worker

// The sandbox interface which enable you to run your program indepandtly of the implementation of the sandbox
type Sandbox interface {
	// Destroy the sandbox on the main system.
	// Usually just after its creation with defer.
	Destroy()

	// Run a command in the sandbox and get its outputs
	// (stdout, stderr, exit_code)
	Run(command string, connection bool) (string, string, int)

	// Get logs linked to the Sandbox
	// Useful for debugging
	GetLogs() string
}
