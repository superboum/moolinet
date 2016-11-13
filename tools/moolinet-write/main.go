package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Need two arguments: " + os.Args[0] + " [path] [content]")
	}
	path := os.Args[1]
	content := os.Args[2]

	f, err := os.Create(path)
	check(err)
	defer f.Close()

	f.WriteString(content)
	check(err)
}
