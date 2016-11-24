package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
	"github.com/superboum/moolinet/lib/persistence"
	"github.com/superboum/moolinet/lib/tools"
	"github.com/superboum/moolinet/lib/web"
)

func main() {
	var config = flag.String("config", "moolinet.json", "Choose a config file for moolinet")

	err := tools.LoadGeneralConfigFromFile(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = persistence.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	judge, err := judge.NewSimpleJudge(&tools.GeneralConfig)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/auth/", web.NewAuthController("/api/auth/"))
	mux.Handle("/api/challenge/", web.NewChallengeController(judge))
	mux.Handle("/api/job/", web.NewJobController(judge, "/api/job/"))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listen on 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
