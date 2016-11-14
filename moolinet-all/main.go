package main

import (
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
	"github.com/superboum/moolinet/lib/tools"
	"github.com/superboum/moolinet/lib/web"
)

func main() {
	judge, err := judge.NewSimpleJudge(&tools.Config{ChallengesPath: "./tests/loadChallengeTest"})
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/challenge/", web.NewChallengeController(judge))
	mux.Handle("/api/job/", web.NewJobController(judge, "/api/job/"))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
