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
	defer func() {
		errDefer := f.Close()
		if errDefer != nil {
			log.Fatal("Unable to close the reader: " + errDefer.Error())
		}
	}()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal("Unable to write content in the file: " + err.Error())
	}

	check(err)
}
