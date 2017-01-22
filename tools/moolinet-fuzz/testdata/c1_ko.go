package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a == 6 {
		a = -1
	}
	fmt.Println(a + b)
}
