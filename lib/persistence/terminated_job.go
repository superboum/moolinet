package persistence

import (
	"time"

	// Used for the driver database/sql
	_ "github.com/mattn/go-sqlite3"
	"github.com/superboum/moolinet/lib/tasks"
)

// TerminatedJob is a struct which allow to save job that ended in base
type TerminatedJob struct {
	UUID      string
	Challenge string
	Username  string
	Status    int
	Created   time.Time
}

// NewTerminatedJobFromJob returns a terminated and persisted job
func NewTerminatedJobFromJob(slug string, user *User, job *tasks.Job) (*TerminatedJob, error) {
	tj := &TerminatedJob{UUID: job.UUID, Username: user.Username, Status: job.Status, Challenge: slug, Created: time.Now()}

	stmt, err := DB.Prepare("INSERT INTO job(uuid, challenge, username, status, created) values(?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(tj.UUID, tj.Challenge, tj.Username, tj.Status, tj.Created)
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	return tj, nil
}

// GetLastNJobs gets a given amount of jobs from the database
func GetLastNJobs(count int) ([]*TerminatedJob, error) {
	res := make([]*TerminatedJob, 0)

	stmt, err := DB.Prepare("SELECT uuid, challenge, username, status, created FROM job ORDER BY created DESC LIMIT 0,?")
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(count)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		t := &TerminatedJob{}
		err := rows.Scan(&t.UUID, &t.Challenge, &t.Username, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}

	return res, nil
}

// GetValidatedChallengePerUser returns a map containing for each user its validated challenges
func GetValidatedChallengePerUser() (map[string][]*TerminatedJob, error) {
	res := make(map[string][]*TerminatedJob)

	rows, err := DB.Query("SELECT uuid, challenge, username, status, created FROM job WHERE status=3")
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		t := &TerminatedJob{}
		err := rows.Scan(&t.UUID, &t.Challenge, &t.Username, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		if _, ok := res[t.Username]; !ok {
			res[t.Username] = make([]*TerminatedJob, 0)
		}
		res[t.Username] = append(res[t.Username], t)
	}

	return res, nil
}
