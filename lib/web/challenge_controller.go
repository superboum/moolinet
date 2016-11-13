package web

import (
	"encoding/json"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
)

type ChallengeController struct {
	judge *judge.Judge
}

func NewChallengeController(j *judge.Judge) *ChallengeController {
	c := new(ChallengeController)
	c.judge = j
	return c
}

func (c *ChallengeController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	encoder.Encode(c.judge.PublicChallenges)
	//fmt.Fprintf(res, "Hi there, I love %s!", req.URL.Path[15:])
}
