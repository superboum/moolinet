package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	content := make([]byte, size)
	for i := 0; i < size; i++ {
		content[i] = 'a'
	}
	fmt.Printf("%s\n", content)
}
