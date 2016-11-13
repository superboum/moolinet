package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
)

type JobController struct {
	judge *judge.Judge
}

type postJob struct {
	Slug string
	Vars map[string]string
}

func NewJobController(j *judge.Judge) *JobController {
	jc := new(JobController)
	jc.judge = j
	return jc
}

func (jc *JobController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		jc.createJob(res, req)
	} else {
	}
}

func (jc *JobController) createJob(res http.ResponseWriter, req *http.Request) {
	newJob := postJob{}

	encoder := json.NewEncoder(res)
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newJob)

	if err != nil {
		res.WriteHeader(400)
		encoder.Encode(APIError{"Your request is malformed", "Put your JSON body in a validator and check the API"})
		return
	}

	job, err := jc.judge.Submit(newJob.Slug, newJob.Vars)
	if err != nil {
		res.WriteHeader(500)
		encoder.Encode(APIError{"Unable to perform your request", "Please contact a server administrator"})
		log.Println(err)
		return
	}
	encoder.Encode(job)
}

func (jc *JobController) getJobStatus(res http.ResponseWriter, req *http.Request) {
}
