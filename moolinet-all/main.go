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

	auth := web.NewAuthMiddleware()

	mux := http.NewServeMux()
	// Authenticated
	mux.Handle("/api/challenge/", auth.CheckAuthentication(web.NewChallengeController(judge)))
	mux.Handle("/api/job/", auth.CheckAuthentication(web.NewJobController(judge, auth, "/api/job/")))
	// Public
	mux.Handle("/api/auth/", web.NewAuthController("/api/auth/", auth))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listening on " + tools.GeneralConfig.ListenAddr)
	err = http.ListenAndServe(tools.GeneralConfig.ListenAddr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
