package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//mux.Handle("/api/challenge", NewChallengeController())
	//mux.Handle("/api/job", NewChallengeController())
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
