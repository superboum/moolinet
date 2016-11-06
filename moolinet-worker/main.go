package main

import (
	"fmt"

	"github.com/superboum/moolinet/lib/worker"
)

func main() {
	s, err := worker.NewDockerSandbox("superboum/moolinet-golang")
	if err != nil {
		panic(err)
	}

	output, err := s.Run(
		[]string{"GOPATH=/home/moolinet"},
		[]string{"go", "get", "github.com/superboum/moolinet"},
		10, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("output-->" + output)

	fmt.Println(s.GetLogs())
}
