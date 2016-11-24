package persistence

import (
	// Used for the driver database/sql
	_ "github.com/mattn/go-sqlite3"
	"github.com/superboum/moolinet/lib/tasks"
)

// TerminatedJob is a struct which allow to save job that ended in base
type TerminatedJob struct {
}

// AddJob adds a new job to the database
func AddJob(job *tasks.Job) error {
	return nil
}

// GetJob gets a given job from database
func GetJob(UUID string) (*tasks.Job, error) {
	return nil, nil
}
