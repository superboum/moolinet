package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
	"github.com/superboum/moolinet/lib/persistence"
	"github.com/superboum/moolinet/lib/tools"
)

// JobController is a controller used to manage Jobs.
type JobController struct {
	judge   *judge.Judge
	auth    *AuthMiddleware
	baseURL string
}

type postJob struct {
	Slug string
	Vars map[string]string
}

type getJobExecution struct {
	Description string
	Output      string `json:",omitempty"`
	Error       string
	Run         bool
}

type getJob struct {
	UUID       string
	Status     int
	Executions []getJobExecution
}

// NewJobController returns a new JobController from a Judge and a baseURL.
// The baseURL is used to clean URLs when generating job IDs.
func NewJobController(j *judge.Judge, a *AuthMiddleware, baseURL string) *JobController {
	jc := new(JobController)
	jc.judge = j
	jc.baseURL = baseURL
	jc.auth = a
	return jc
}

// ServeHTTP is a router for the Job Controller:
//
// - If the method is POST, it creates a new Job ;
// - Otherwise, if the URL is not malformed, it returns the status of a Job
//
// The expected URL is {baseURL}{jobUUID}.
func (jc *JobController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	action := req.URL.Path[len(jc.baseURL):]
	if req.Method == "POST" {
		jc.createJob(res, req)
	} else if action == "ranking" {
		// @FIXME it would be a better idea to create a ranking controller
		jc.getRanking(res, req)
	} else if len(action) > 0 {
		jc.getJobStatus(res, req)
	}
}

func checkEncode(errEncode error) {
	if errEncode != nil {
		log.Println("There was an error while encoding the response: " + errEncode.Error())
	}
}

func (jc *JobController) getRanking(res http.ResponseWriter, req *http.Request) {
	list, err := persistence.GetValidatedChallengePerUser()
	encoder := json.NewEncoder(res)
	if err != nil {
		res.WriteHeader(500)
		checkEncode(encoder.Encode(APIError{"Unable to perform your request", "Please contact a server administrator"}))
		log.Println(err)
		return
	}
	checkEncode(encoder.Encode(list))
}

func (jc *JobController) createJob(res http.ResponseWriter, req *http.Request) {
	newJob := postJob{}

	encoder := json.NewEncoder(res)
	decoder := json.NewDecoder(io.LimitReader(req.Body, tools.GeneralConfig.MaxSubmissionSize))
	err := decoder.Decode(&newJob)

	if err != nil {
		res.WriteHeader(400)
		checkEncode(encoder.Encode(APIError{"Your request is malformed", "Put your JSON body in a validator and check the API"}))
		return
	}

	u, err := jc.auth.GetUser(req)
	if err != nil {
		res.WriteHeader(500)
		checkEncode(encoder.Encode(APIError{"Unable to perform your request", "Please contact a server administrator"}))
		log.Println(err)
		return
	}

	job, err := jc.judge.Submit(newJob.Slug, newJob.Vars, u)
	if err != nil {
		res.WriteHeader(500)
		checkEncode(encoder.Encode(APIError{"Unable to perform your request", "Please contact a server administrator"}))
		log.Println(err)
		return
	}
	checkEncode(encoder.Encode(job))
}

func (jc *JobController) getJobStatus(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	UUID := req.URL.Path[len(jc.baseURL):]
	job, err := jc.judge.GetJob(UUID)
	if err != nil {
		res.WriteHeader(404)
		checkEncode(encoder.Encode(APIError{"The UUID was not found", "Check that your job exists"}))
		return
	}

	pjob := getJob{
		UUID:       job.UUID,
		Status:     job.Status,
		Executions: make([]getJobExecution, len(job.Executions)),
	}
	for i, e := range job.Executions {
		pjob.Executions[i].Description = e.Description
		pjob.Executions[i].Error = e.Error
		pjob.Executions[i].Run = e.Run
		if e.Public {
			pjob.Executions[i].Output = e.Output
		}
	}

	checkEncode(encoder.Encode(pjob))
}
