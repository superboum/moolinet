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
	if err != nil {
		panic(err)
	}
	fmt.Println("output-->" + output)

	output, err = s.Run([]string{"go", "get", "-d", "github.com/superboum/atuin/..."}, 120, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("output-->" + output)

	output, err = s.Run([]string{"go", "install", "github.com/superboum/atuin/..."}, 120, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("output-->" + output)

	output, err = s.Run([]string{"atuin-front"}, 30, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("output-->" + output)

	fmt.Println(s.GetLogs())
}
