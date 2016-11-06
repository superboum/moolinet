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

	output, err := s.Run([]string{"go", "get", "github.com/superboum/atuin"}, 120, true)
	fmt.Println("output-->" + output)
	if err != nil {
		panic(err)
	}

	output, err = s.Run([]string{"go", "get", "-d", "github.com/superboum/atuin/..."}, 120, true)
	fmt.Println("output-->" + output)
	if err != nil {
		panic(err)
	}

	output, err = s.Run([]string{"go", "install", "github.com/superboum/atuin/..."}, 120, false)
	fmt.Println("output-->" + output)
	if err != nil {
		panic(err)
	}

	output, err = s.Run([]string{"atuin-front"}, 30, false)
	fmt.Println("output-->" + output)
	if err != nil {
		panic(err)
	}

	fmt.Println(s.GetLogs())
}
