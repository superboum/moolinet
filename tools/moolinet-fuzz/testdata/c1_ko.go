package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if b == 42 {
		b = -1
	}
	fmt.Println(a+1, b+1)
}
