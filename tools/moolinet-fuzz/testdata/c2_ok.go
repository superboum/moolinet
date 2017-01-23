package main

import "fmt"

func main() {
	var n, a, b, c int64
	var op string

	// Read number of test cases
	fmt.Scan(&n)

	// Execute operations
	for i := int64(0); i < n; i++ {
		fmt.Scan(&a, &op, &b)
		switch op {
		case "ADD":
			c = a + b
		case "SUB":
			c = a - b
		case "MUL":
			c = a * b
		}
		fmt.Println(c)
	}
}
