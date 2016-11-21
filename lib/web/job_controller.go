package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
)

// JobController is a controller used to manage Jobs.
type JobController struct {
	judge   *judge.Judge
	baseURL string
}

type postJob struct {
	Slug string
	Vars map[string]string
}

// NewJobController returns a new JobController from a Judge and a baseURL.
// The baseURL is used to clean URLs when generating job IDs.
func NewJobController(j *judge.Judge, baseURL string) *JobController {
	jc := new(JobController)
	jc.judge = j
	jc.baseURL = baseURL
	return jc
}

// ServeHTTP is a router for the Job Controller:
//
// - If the method is POST, it creates a new Job ;
// - Otherwise, if the URL is not malformed, it returns the status of a Job
//
// The expected URL is {baseURL}{jobUUID}.
func (jc *JobController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		jc.createJob(res, req)
	} else if len(req.URL.Path[len(jc.baseURL):]) > 0 {
		jc.getJobStatus(res, req)
	}
}

func checkEncode(errEncode error) {
	if errEncode != nil {
		log.Println("There was an error while encoding the response: " + errEncode.Error())
	}
}

func (jc *JobController) createJob(res http.ResponseWriter, req *http.Request) {
	newJob := postJob{}

	encoder := json.NewEncoder(res)
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newJob)

	if err != nil {
		res.WriteHeader(400)
		checkEncode(encoder.Encode(APIError{"Your request is malformed", "Put your JSON body in a validator and check the API"}))
		return
	}

	job, err := jc.judge.Submit(newJob.Slug, newJob.Vars)
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

	checkEncode(encoder.Encode(job))
}
