package main

import (
	"fmt"

	"github.com/superboum/moolinet/lib/worker"
)

func main() {
	s, err := worker.NewDockerSandbox("golang:alpine")
	if err != nil {
		panic(err)
	}
	fmt.Println(s.GetLogs())
}
