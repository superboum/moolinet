package main

import "fmt"

func main() {
	var size int
	for {
		nb, err := fmt.Scanf("%d", &size)
		if nb == 0 || err != nil {
			break
		}

		content := make([]byte, size)
		for i := 0; i < size; i++ {
			content[i] = 'a'
		}
		fmt.Printf("%s\n", content)
	}
}
